# Producer Web Service 

This is a basic webservice which applications can send their logs via an http request, then the web service
will appropriately store the log into kafka where it should go.

TO USE:

From the main project directory run:

`sh bin/start\_kafka.sh

`sh bin/run\_collector.sh


Then you can use the test client in the test-client folder:

`go run log-collector-client.go
