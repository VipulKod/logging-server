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

