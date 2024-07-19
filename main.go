package main

import (
	"log"
	"net/http"

	dbconnection "loggingserver/db"
	handlers "loggingserver/handlers"
	queues "loggingserver/queues"
	serverWorker "loggingserver/workers"
)

func main() {
	db, err := dbconnection.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}

	queue := queues.NewQueue()

	doneChans := make([]chan bool, 3)
	for i := 0; i < 3; i++ {
		doneChans[i] = make(chan bool, 1)
	}

	go serverWorker.Worker(queue.InfoQueue, db, doneChans[0])
	go serverWorker.Worker(queue.WarnQueue, db, doneChans[1])
	go serverWorker.Worker(queue.ErrorQueue, db, doneChans[2])

	http.Handle("/log", handlers.LogHandler(db, queue))
	http.Handle("/logs", handlers.LogsHandler(db))

	log.Fatal(http.ListenAndServe(":8080", nil))

	// Wait for all workers to finish
	for _, ch := range doneChans {
		<-ch
	}
}
