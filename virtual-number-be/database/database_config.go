package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func DatabaseConnection() (*sql.DB, error) {

	DBUser := os.Getenv("DBUser")
	DBPass := os.Getenv("DBPass")
	DBHost := os.Getenv("DBHost")
	DBPort := os.Getenv("DBPort")
	DBName := os.Getenv("DBName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DBUser, DBPass, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	// Check if the database connection is established
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connecting to database ...")
	time.Sleep(time.Second)
	fmt.Println("Connected to database successfully")

	db.SetConnMaxIdleTime(1 * time.Hour)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)

	return db, nil
}
