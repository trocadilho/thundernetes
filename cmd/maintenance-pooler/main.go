package main

import (
	"bytes"
	"fmt"
	"encoding/json"
	"time"
	"net/http"
)

func main() {
	for {
		scheduledEvent, err := GetScheduledEvents()
		if (err != nil) {
			fmt.Println("error get")
		}
		fmt.Println(scheduledEvent.DocumentIncarnation)
	
		events := scheduledEvent.Events
	
		if (len(events) > 0) {
			statusCode, err := ConfirmScheduledEvents(scheduledEvent.Events[0].EventID)
			if (err != nil) {
				fmt.Println("error confirm")
			}
			fmt.Println(statusCode)
		}
		time.Sleep(1 * time.Second)
	}
}

func GetScheduledEvents() (ScheduledEvent, error) {
	scheduledEvent := ScheduledEvent {}
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://169.254.169.254/metadata/scheduledevents", nil)
	if err != nil {
		return scheduledEvent, err
	}

	q := req.URL.Query()
	q.Add("api-version", "2020-07-01")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Metadata", "true")

	res, err := client.Do(req)
	if err != nil {
		return scheduledEvent, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&scheduledEvent)

	return scheduledEvent, err
}

func ConfirmScheduledEvents(eventId string) (int, error) {
	statusCode := 0
	events := ConfirmScheduledEvent{
		StartRequests: []StartRequest{{EventID: eventId}},
	}
	client := http.Client{}
	postBody, _ := json.Marshal(events)
	buffer := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", "http://169.254.169.254/metadata/scheduledevents", buffer)
	if err != nil {
		return statusCode, err
	}

	q := req.URL.Query()
	q.Add("api-version", "2020-07-01")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Metadata", "true")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	return res.StatusCode, err
}