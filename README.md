# Producer Web Service 

###To be used with: https://github.com/kwilsonO/log-daemon and https://github.com/kwilsonO/cassandra-consumer

This is a basic webservice which applications can send their logs via an http request, then the web service
will appropriately store the log into kafka where it should go.

TO USE:

From the main project directory run:

	sh bin/start_kafka.sh

	sh bin/run_collector.sh


Then you can use the test client in the test-client folder:

	go run log-collector-client.go
