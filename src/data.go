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
		DBName:               "taaf",
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
	fmt.Println("database connected")
	check_and_seed_table(db)
	fmt.Println("database seeded")
	return db
}

func check_and_seed_table(db *sql.DB) error {
	db_query := `CREATE DATABASE IF NOT EXISTS taaf`
	fmt.Println("database created")
	query_db(db, db_query)
	use_query := `USE taaf`
	fmt.Println("on TAAF DB")
	query_db(db, use_query)
	query := `CREATE TABLE IF NOT EXISTS video(
		video_id int primary key auto_increment, 
		target_resolution text, 
		source_location text,
		transcode_progress INT DEFAULT 0,
		created_at datetime default CURRENT_TIMESTAMP, 
		updated_at datetime default CURRENT_TIMESTAMP
	)`
	_, err := query_db(db, query)
	fmt.Println("table created")
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	fmt.Println("table created without error")
	return nil
}

func query_db(db *sql.DB, query string) (*sql.Rows, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when running query", err)
		return nil, err
	}
	return res, nil
}

func QueryDB(query string) (*sql.Rows, error) {
	return query_db(db, query)
}
