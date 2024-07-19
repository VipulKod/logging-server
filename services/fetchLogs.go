package services

import (
	"database/sql"
	"time"

	logging "log"

	models "loggingserver/models"
)

// FetchLogs retrieves logs from the database based on provided query parameters.
// It returns a slice of logs matching the query criteria.
func FetchLogs(params models.LogQueryParams, db *sql.DB) ([]models.LogResponse, error) {
	var logs []models.LogResponse

	// Base query
	query := `
		SELECT id, severity, serviceName, message, timestamp FROM logs 
		WHERE (timestamp >= $1 OR $1 IS NULL)
		AND (timestamp <= $2 OR $2 IS NULL)
		AND (severity = $3 OR $3 = '')
		AND (serviceName = $4 OR $4 = '')
	`

	// Prepare optional parameters
	var startTimestamp, endTimestamp *time.Time
	if params.StartTimestamp != "" {
		start, err := time.Parse(time.RFC3339, params.StartTimestamp)
		if err != nil {
			logging.Printf("Error parsing start timestamp: %v\n", err)
			return nil, err
		}
		startTimestamp = &start
	}
	if params.EndTimestamp != "" {
		end, err := time.Parse(time.RFC3339, params.EndTimestamp)
		if err != nil {
			logging.Printf("Error parsing end timestamp: %v\n", err)
			return nil, err
		}
		endTimestamp = &end
	}

	// Execute the query with provided parameters
	rows, err := db.Query(query, startTimestamp, endTimestamp, params.Severity, params.ServiceName)
	if err != nil {
		logging.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and populate logs slice
	for rows.Next() {
		var log models.LogResponse
		err := rows.Scan(&log.Id, &log.Severity, &log.ServiceName, &log.Message, &log.Timestamp)
		if err != nil {
			logging.Printf("Error scanning row: %v\n", err)
			continue
		}
		logs = append(logs, log)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		logging.Printf("Rows iteration error: %v\n", err)
	}

	return logs, nil
}
