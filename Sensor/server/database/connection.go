package database

import (
	"database/sql"
	"fmt"
	"server/logger"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "990527"
	dbname   = "Database"
)

func setupDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("database connected")
	return db, nil
}

var Database *sql.DB

func init() {
	db, err := setupDB()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return
	}
	Database = db
}
