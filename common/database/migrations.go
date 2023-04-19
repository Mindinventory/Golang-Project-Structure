package database

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var filenameRegexp = regexp.MustCompile(`V(\d+)__(.*).sql$`)

// Migration is a Flyway migration file read from disk.
type Migration struct {
	Path string
	Data []byte
}

// Migrations is a sortable collection of migrations.
type Migrations []*Migration

func (s Migrations) Len() int      { return len(s) }
func (s Migrations) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Migrations) Less(i, j int) bool {
	return s[i].Version() < s[j].Version()
}

// Filename is the base name of the migration path with the directory removed.
func (m *Migration) Filename() string {
	return filepath.Base(m.Path)
}

// Name is the name of the migration.
func (m *Migration) Name() string {
	matches := filenameRegexp.FindStringSubmatch(m.Filename())

	if len(matches) < 3 {
		return ""
	}
	return matches[2]
}

// Version is the sortable number of the migration.
func (m *Migration) Version() int {
	name := m.Filename()

	matches := filenameRegexp.FindStringSubmatch(name)

	if len(matches) < 3 {
		return -1
	}

	i, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1
	}

	return i
}

// SQL returns the string SQL for passing to (*sql.DB).Exec().
func (m *Migration) SQL() string {
	return string(m.Data)
}

func LoadMigrations(path string) (Migrations, error) {
	var migrations []*Migration

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		migrations = append(migrations, &Migration{
			Path: path,
			Data: b,
		})

		return nil
	})

	return migrations, err
}
