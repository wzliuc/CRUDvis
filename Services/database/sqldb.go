package database

import (
	"database/sql"
	"webservice/handlers/logger"
)

// Db represents the database conenction
var Db *sql.DB

// SetupDataBase initialises the database connection
func SetupDataBase() {
	logger.LogInfo("Initialising database...")
	var err error
	Db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/thegreatwall")
	if err != nil {
		logger.LogErr(err)
	}
	logger.LogInfo("Database initialisation complete.")
}
