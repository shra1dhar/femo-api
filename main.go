package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"time"
)

// https://medium.com/geekculture/work-with-go-postgresql-using-pgx-caee4573672

const (
	DbHost = "db"
	DbUser = "postgres-dev"
	DbPassword = "mysecretpassword"
	Dbname  = "dev"
	Migration = `CREATE TABLE IF NOT EXISTS bulletins
id serial PRIMARY KEY
author TEXT NOT NULL
content TEXT NOT NULL
created_at timestamp with timezone DEFAULT current_timestamp`
)

type Bulletin struct {
	Author string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

var db *sql.DB

func GetBulletins() ([]Bulletin, error) {
	return nil, nil
}

func AddBulletin(bulletin Bulletin) error {
	return nil
}

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
