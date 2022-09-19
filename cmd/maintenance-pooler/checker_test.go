package main

import "testing"

type MockMetadataClient struct {
	mockedScheduledEvent ScheduledEvent
}

func (c MockMetadataClient) GetScheduledEvents() (ScheduledEvent, error) {
	return c.mockedScheduledEvent, nil
}

func (c MockMetadataClient) ConfirmScheduledEvent(eventId string) (statusCode int, err error) {
	return 200, nil
}

type MockMaintenanceNotifier struct{}

func (n MockMaintenanceNotifier) Notify() error {
	return nil
}

func Test(t *testing.T) {
	t.Run("abc", func(t *testing.T) {
		client := MockMetadataClient{}
		notifier := MockMaintenanceNotifier{}
		checker := NewChecker(client, notifier)

		err := checker.Check()

		if err != nil {
			t.Fatalf("not expected")
		}
	})
}
