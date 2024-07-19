package queues

import (
	models "loggingserver/models"
)

type Queue struct {
	InfoQueue  chan models.LogRequest
	WarnQueue  chan models.LogRequest
	ErrorQueue chan models.LogRequest
}

func NewQueue() *Queue {
	return &Queue{
		InfoQueue:  make(chan models.LogRequest),
		WarnQueue:  make(chan models.LogRequest),
		ErrorQueue: make(chan models.LogRequest),
	}
}
