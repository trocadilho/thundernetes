package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func main() {
	scheduledEvent, err := getScheduledEvents()
	if (err != nil) {
		fmt.Println("error")
	}
	fmt.Println(scheduledEvent.DocumentIncarnation)
}

func getScheduledEvents() (ScheduledEvent, error) {
	scheduledEvent := ScheduledEvent {}
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://169.254.169.254/metadata/scheduledevents", nil)
	if err != nil {
		fmt.Println("error 1")
		return scheduledEvent, err
	}

	q := req.URL.Query()
	q.Add("api-version", "2020-07-01")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Metadata", "true")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error 2")
		return scheduledEvent, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&scheduledEvent)

	return scheduledEvent, err
}