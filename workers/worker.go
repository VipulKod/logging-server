package workers

import (
	"database/sql"

	models "loggingserver/models"
	service "loggingserver/services"
)

// Worker processes incoming log requests from a channel,
// saves them to the database, and signals completion through a channel.
func Worker(queue chan models.LogRequest, db *sql.DB, doneChan chan bool) {
	defer close(doneChan)
	for log := range queue {
		service.SaveLog(log, db)
		doneChan <- true // Signal that processing for this log is complete
	}
}
