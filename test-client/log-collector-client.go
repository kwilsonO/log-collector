package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type KafkaMsg struct {
	Topic string
	Key   string
	Value string
}

func main() {

	var kmsg KafkaMsg = KafkaMsg{
		Topic: "important",
		Key:   "TestKey",
		Value: "my test value",
	}

	b, err := json.Marshal(kmsg)
	if err != nil {
		fmt.Println("Could not marshal json")
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(b))
	//req, err := http.NewRequest("GET", "http://example.com/?data", nil)
	if err != nil {
		fmt.Printf("FAILED !!! %s", err)
	}

	client := &http.Client{}

	_, err = client.Do(req)

	if err != nil {
		fmt.Printf("FAILED EVEN MORE !!! %s", err)

	}

}
