package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DataBase() {
	db = create_and_return_db()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func create_and_return_db() (db *sql.DB) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "this-just-is-not-good-practice",
		Net:                  "tcp",
		Addr:                 "mysql-taaf:3306",
		DBName:               "",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	check_and_seed_table(db)
	return db
}

func check_and_seed_table(db *sql.DB) error {
	db_query := `CREATE DATABASE IF NOT EXISTS taaf`
	query_db(db, db_query)
	use_query := `USE taaf`
	query_db(db, use_query)
	query := `CREATE TABLE IF NOT EXISTS video(
		video_id int primary key auto_increment, 
		target_resolution text, 
		source_location text,
		transcode_progress INT DEFAULT 0,
		created_at datetime default CURRENT_TIMESTAMP, 
		updated_at datetime default CURRENT_TIMESTAMP
	)`
	rows, err := query_db(db, query)
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func query_db(db *sql.DB, query string) (int64, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating rooms table", err)
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return 0, err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return rows, nil
}
