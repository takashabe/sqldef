package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/k0kubun/schemasql/schema"
)

func buildMysqlDSN() string {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = ""
	config.Net = "tcp"
	config.Addr = "127.0.0.1:3306"
	config.DBName = "test"
	return config.FormatDSN()
}

func runMySQLDDL() {
	dsn := buildMysqlDSN()
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	transaction, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	sql := `
		CREATE TABLE user2 (
		  id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
		  name VARCHAR(191) UNIQUE,
		  salt VARCHAR(20),
		  password VARCHAR(40),
		  display_name TEXT,
		  avatar_icon TEXT,
		  created_at DATETIME NOT NULL
		) Engine=InnoDB DEFAULT CHARSET=utf8mb4;
	`

	if _, err := transaction.Exec(sql); err != nil {
		transaction.Rollback()
		log.Fatal(err)
	}
	transaction.Commit()
}

func parseTable() {
	dsn := buildMysqlDSN()
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var table string
	var ddl string
	err = conn.QueryRow("show create table user;").Scan(&table, &ddl)
	if err != nil {
		log.Fatal(err)
	}

	var ddl2 string
	err = conn.QueryRow("show create table user2;").Scan(&table, &ddl2)
	if err != nil {
		log.Fatal(err)
	}

	schema.ParseDDLs(fmt.Sprintf("-- hello\n%s;\n%s;", ddl, ddl2))
}
