package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gopkg.in/testfixtures.v2"
	"log"
	"os"
)

func main() {
	e := godotenv.Load("../.env")
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	conn, err := sql.Open("mysql", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Asia%2FKolkata")
	//conn, err := sql.Open("mysql", "root:mind@/zoomin_dev")
	if err != nil {
		fmt.Println("connection failed")
	} else {
		fmt.Println("connected")
	}
	fixtures, err := testfixtures.NewFiles(conn, &testfixtures.MySQL{},
		"fixtures/user.yml",
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}

}
