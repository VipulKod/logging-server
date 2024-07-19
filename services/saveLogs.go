package services

import (
	"database/sql"
	logging "log"

	models "loggingserver/models"
)

// saveLog saves a single log entry to the database.
// It logs an error if saving fails and returns the error.
func SaveLog(log models.LogRequest, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO logs(severity, serviceName, message) VALUES($1, $2, $3)",
		log.Severity, log.ServiceName, log.Message)
	if err != nil {
		logging.Printf("Error saving log: %v", err)
		return err
	}
	return nil
}
