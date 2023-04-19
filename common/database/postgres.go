package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/avast/retry-go/v3"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DbConfig struct {
	User     string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	Host     string `required:"true" split_words:"true"`
	Port     int    `required:"true" split_words:"true"`
	Name     string `required:"true" split_words:"true"`

	disableSSL bool
}

type RowsAndCount struct {
	c   int
	id  *uuid.UUID
	err error
}

var (
	pool  *pgxpool.Pool
	ponce sync.Once
)

func (dbConfig *DbConfig) URL() string {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	if dbConfig.disableSSL {
		url = fmt.Sprintf("%s?sslmode=disable", url)
	}

	return url
}

func getConfig(dbConfig DbConfig) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(dbConfig.URL())
	if err != nil {
		return nil, err
	}

	config.ConnConfig.LogLevel = pgx.LogLevelDebug

	return config, nil
}

func resetPool() {
	pool = nil
	ponce = sync.Once{}
}

func CreateConnectionPoolWithConnectionLimit(dbConfig DbConfig, maxConnections int) (*pgxpool.Pool, error) {
	config, err := getConfig(dbConfig)
	if err != nil {
		return nil, err
	}
	config.MaxConns = int32(maxConnections)
	return createConnectionPool(config)
}

func CreateConnectionPool(dbConfig DbConfig) (*pgxpool.Pool, error) {
	config, err := getConfig(dbConfig)
	if err != nil {
		return nil, err
	}
	// well make our default pool size 10. The pacakge default is 4... ouch!
	config.MaxConns = int32(10)
	return createConnectionPool(config)
}

func createConnectionPool(config *pgxpool.Config) (*pgxpool.Pool, error) {
	var err error
	ponce.Do(func() {
		err = retry.Do(func() error {
			p, err := pgxpool.ConnectConfig(context.Background(), config)
			if err != nil {
				return err
			}
			pool = p
			return nil
		},
			retry.Delay(1*time.Second))
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return pool, err
}

func GetConnectionPool() (*pgxpool.Pool, error) {
	if pool == nil {
		return nil, errors.New("connection Pool has not been created")
	}
	return pool, nil
}

func ClosePool() error {
	if pool == nil {
		return errors.New("connection Pool has not been created")
	}
	pool.Close()
	return nil
}

func CountRowsAndLastId(ctx context.Context, table string, descending bool, ch chan RowsAndCount) error {
	if pool == nil {
		return errors.New("connection Pool has not been created")
	}
	idc := table[:len(table)-1]

	batch := pgx.Batch{}
	batch.Queue(fmt.Sprintf("select count(%vid) from assets.%v", idc, table))
	if descending {
		batch.Queue(fmt.Sprintf("select %vid from assets.%v order by %vid asc limit 1", idc, table, idc))
	} else {
		batch.Queue(fmt.Sprintf("select %vid from assets.%v order by %vid desc limit 1", idc, table, idc))
	}

	con, err := pool.Acquire(ctx)
	if err != nil {
		ch <- RowsAndCount{c: 0, id: nil, err: err}
		return nil
	}
	defer con.Release()

	br := con.SendBatch(ctx, &batch)

	c := 0
	err = br.QueryRow().Scan(&c)

	switch err {
	case pgx.ErrNoRows:
		ch <- RowsAndCount{c: 0, id: nil, err: nil}
		return nil
	case nil:
	default:
		ch <- RowsAndCount{c: 0, id: nil, err: err}
		return nil
	}

	u := new(uuid.UUID)
	err = br.QueryRow().Scan(u)

	switch err {
	case pgx.ErrNoRows:
		ch <- RowsAndCount{c: 0, id: nil, err: nil}
		return nil
	case nil:
	default:
		ch <- RowsAndCount{c: 0, id: nil, err: err}
		return nil
	}

	ch <- RowsAndCount{c: c, id: u, err: nil}
	return nil
}

func InsertBatch(ctx context.Context, batch *pgx.Batch) error {
	return ExecuteBatch(ctx, batch, "inserted", true)
}

func UpdateBatch(ctx context.Context, batch *pgx.Batch) error {
	return ExecuteBatch(ctx, batch, "updated", true)
}

func DeleteBatch(ctx context.Context, batch *pgx.Batch) error {
	return ExecuteBatch(ctx, batch, "deleted", false)
}

func ExecuteBatch(ctx context.Context, batch *pgx.Batch, action string, verify bool) error {
	if pool == nil {
		return errors.New("connection Pool has not been created")
	}
	c, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer c.Release()

	tx, err := c.Begin(ctx)
	if err != nil {
		return err
	}

	br := tx.SendBatch(ctx, batch)

	for i := 0; i < batch.Len(); i++ {
		ct, err := br.Exec()
		if err != nil {

			if er := tx.Rollback(ctx); er != nil {
			}

			return err
		}

		if ct.RowsAffected() != 1 {
			msg := fmt.Sprintf("no row %v", action)

			if verify {
				if err := br.Close(); err != nil {
				}
				if er := tx.Rollback(ctx); er != nil {
				}

				return errors.New(msg)
			}
		}
	}

	if err := br.Close(); err != nil {
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// https://www.starkandwayne.com/blog/uuid-primary-keys-in-postgresql/
// https://wiki.postgresql.org/images/3/35/Pagination_Done_the_PostgreSQL_Way.pdf
