# logging-server
A simple logging server build with GoLang and other standard Go libraries

# Prerequisites: 
* Ensure you have Go installed on your machine. You can download it from the official Go website.
* Clone the Repository: Clone the project repository to your local machine using Git.
* Navigate to the Project Directory: Change your current directory to the root of the cloned project.
* Install Dependencies: Run go mod tidy to automatically download and install dependencies defined in the go.mod file.
* Configure Database: Update the connection string in db/connection.go with your PostgreSQL credentials and details. Ensure PostgreSQL is running on your machine or accessible via the network.
* Start the Application: Execute go run main.go to start the application. The server will listen on port 8080.
* Testing: Test the functionality by sending log entries through the /log endpoint and querying them through the /logs endpoint.

# Architecture Overview
* Log storage can be challenge sometimes especially when there are lot of logs that needs to be collected concurrently from various sources to know the exact reason of what can go wrong or who changed what or to create a modern dashboard for viewing logs or stats. For such use cases a logging application that stores all the necessary logs from various other sources like microservies can be of great use.
* For this use case i have used go programming language to build a mock solution for scalable logging application. Its still in its initial stage and this concept can be leveraged using dedicated pub sub model based applications such as kafka, rabbit mq etc to have a scalable solution.
* For this project I have used dedicated Goroutines for Log Collection: Within each component, a dedicated goroutine collects logs and sends them to a central buffer via channels. This ensures that log collection does not hinder the component's primary functions.
* Central Buffer: Implemented as a channel, the central buffer receives logs from all components. This buffer temporarily holds log messages, preventing any loss during transmission.
* Also there is a sepration of queue based on the priority levels such as "INFO", "WARN", "ERROR" and each has different priority based on the urgency.
* Continuous Monitoring and Processing: Additional goroutines monitor the central buffer. Upon reaching a predetermined threshold or after a specific interval, these goroutines process the logs (such as aggregating, filtering) and then forward them to a persistent storage system or a log analysis tool.
* Persistent Storage or Analysis Tool: The processed logs are finally stored in a database.
