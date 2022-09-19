package repository

import (
	"database/sql"
	"server/database"
	"server/logger"
)

func checkRowExistbyId(query string, args ...interface{}) error {
	var rowExists bool
	err := database.Database.QueryRow(query, args...).Scan(&rowExists)
	if err == sql.ErrNoRows {
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	return nil
}
