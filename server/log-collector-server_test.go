package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// In normal operation, we expect one access log entry,
// and one data collector entry. Let's assume both will succeed.
// We should return a HTTP 200 status.
func TestCollectSuccessfully(t *testing.T) {
	dataCollectorMock := mocks.NewSyncProducer(t, nil)
	dataCollectorMock.ExpectSendMessageAndSucceed()

	// Now, use dependency injection to use the mocks.
	s := &Server{
		DataCollector: dataCollectorMock,
	}

	// The Server's Close call is important; it will call Close on
	// the two mock producers, which will then validate whether all
	// expectations are resolved.
	defer safeClose(t, s)

	var kmsg KafkaMsg = KafkaMsg{
		Topic: "TestTopic",
		Key:   "TestKey",
		Value: "my test value",
	}

	b, err := json.Marshal(kmsg)
	if err != nil {
		fmt.Println("Could not marshal json")
		return
	}

	req, err := http.NewRequest("POST", "http://localhost/", bytes.NewBuffer(b))
	//req, err := http.NewRequest("GET", "http://example.com/?data", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	s.Handler().ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("Expected HTTP status 200, found %d", res.Code)
	}

	if string(res.Body.Bytes()) != "Your data is stored with unique identifier important/0/1" {
		t.Error("Unexpected response body", res.Body)
	}
}

// We don't expect any data collector calls because the path is wrong,
// so we are not setting any expectations on the dataCollectorMock. It
// will still generate an access log entry though.
func TestWrongPath(t *testing.T) {
	dataCollectorMock := mocks.NewSyncProducer(t, nil)

	s := &Server{
		DataCollector: dataCollectorMock,
	}
	defer safeClose(t, s)

	req, err := http.NewRequest("GET", "http://example.com/wrong?data", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	s.Handler().ServeHTTP(res, req)

	if res.Code != 404 {
		t.Errorf("Expected HTTP status 404, found %d", res.Code)
	}
}

func safeClose(t *testing.T, o io.Closer) {
	if err := o.Close(); err != nil {
		t.Error(err)
	}
}
