package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
)

var connStr string = "tcp://localhost:9000"

func write(ch *sql.DB) error {
	tx, err := ch.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO not_exists (
			ts,
			siteID,
			user.id,
			user.accountID,
			segmentID
		)
		VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for i := 0; i < 3; i++ {
		_, err = stmt.Exec(
			time.Now(),
			1,
			[]string{"a"},
			[]string{"a"},
			1,
		)

		if err != nil {
			stmt.Close()
			tx.Rollback()
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func main() {
	db, err := sql.Open("clickhouse", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(5 * time.Second)

	// will fail.
	_, err = db.Query("select * from not_exists")
	if err != nil {
		log.Println(err)
	}

	// will fail.
	err = write(db)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(30 * time.Second)
	// users table exists.
	rows, err := db.Query("select * from visitor_membership")
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()
}
