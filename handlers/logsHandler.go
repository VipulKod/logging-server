package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	models "loggingserver/models"
	services "loggingserver/services"
	utils "loggingserver/utils"
)

func LogsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is accepted", http.StatusMethodNotAllowed)
			return
		}

		var queryParams models.LogQueryParams
		err := utils.ParseQueryParams(r.URL.Query(), &queryParams)
		if err != nil {
			http.Error(w, "Failed to parse query parameters", http.StatusBadRequest)
			return
		}

		logs, err := services.FetchLogs(queryParams, db)
		if err != nil {
			http.Error(w, "Failed to fetch logs", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonResponse, err := json.Marshal(logs)
		if err != nil {
			http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)
	}
}
