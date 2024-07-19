package utils

import (
	"fmt"
	"time"

	models "loggingserver/models"
)

func ParseQueryParams(queryValues map[string][]string, queryParams *models.LogQueryParams) error {
	// StartTimestamp
	if startTimestamps, ok := queryValues["startTimestamp"]; ok && len(startTimestamps) > 0 {
		startTimestamp, err := time.Parse(time.RFC3339, startTimestamps[0])
		if err != nil {
			return fmt.Errorf("failed to parse startTimestamp: %w", err)
		}
		queryParams.StartTimestamp = startTimestamp.Format(time.RFC3339)
	}

	// EndTimestamp
	if endTimestamps, ok := queryValues["endTimestamp"]; ok && len(endTimestamps) > 0 {
		endTimestamp, err := time.Parse(time.RFC3339, endTimestamps[0])
		if err != nil {
			return fmt.Errorf("failed to parse endTimestamp: %w", err)
		}
		queryParams.EndTimestamp = endTimestamp.Format(time.RFC3339)
	}

	// Severity
	if severities, ok := queryValues["severity"]; ok && len(severities) > 0 {
		queryParams.Severity = severities[0]
	}

	// ServiceName
	if serviceNames, ok := queryValues["serviceName"]; ok && len(serviceNames) > 0 {
		queryParams.ServiceName = serviceNames[0]
	}

	return nil
}
