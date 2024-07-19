package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	models "loggingserver/models"
	queues "loggingserver/queues"
)

func LogHandler(db *sql.DB, queue *queues.Queue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure only POST requests are handled
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
			return
		}

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Decode JSON body into LogRequest struct
		var log models.LogRequest
		if err := json.Unmarshal(body, &log); err != nil {
			http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
			return
		}

		// Route the log based on severity
		switch log.Severity {
		case "INFO":
			queue.InfoQueue <- log
		case "WARN":
			queue.WarnQueue <- log
		case "ERROR":
			queue.ErrorQueue <- log
		default:
			http.Error(w, "Invalid severity level", http.StatusBadRequest)
			return
		}

		// Respond with a success message
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "success", "message": "Log received successfully"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(jsonResponse)
	}
}
