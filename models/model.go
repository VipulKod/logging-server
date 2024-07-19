package models

import "time"

type LogRequest struct {
	Severity    string `json:"severity"`
	ServiceName string `json:"serviceName"`
	Message     string `json:"message"`
}

type LogQueryParams struct {
	StartTimestamp string `query:"startTimestamp"`
	EndTimestamp   string `query:"endTimestamp"`
	Severity       string `query:"severity"`
	ServiceName    string `query:"serviceName"`
}

type LogResponse struct {
	Id          int       `json:"id"`
	Severity    string    `json:"severity"`
	ServiceName string    `json:"serviceName"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
}
