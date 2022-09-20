package main

import (
	"context"
	"fmt"
	"time"
)

type Checker struct {
	client   MetadataClient
	notifier MaintenanceNotifier
}

func NewChecker(client MetadataClient, notifier MaintenanceNotifier) *Checker {
	checker := new(Checker)
	checker.client = client
	checker.notifier = notifier
	return checker
}

func (c Checker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := c.Check(ctx)

			if err != nil {
				fmt.Println(err.Error())
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (c Checker) Check(ctx context.Context) error {
	scheduledEvent, err := c.client.GetScheduledEvents()
	if err != nil {
		return err
	}
	fmt.Println(scheduledEvent.DocumentIncarnation)

	for _, event := range scheduledEvent.Events {
		err = c.notifier.Notify(ctx)
		if err != nil {
			return err
		}

		statusCode, err := c.client.ConfirmScheduledEvent(event.EventID)
		fmt.Println(statusCode)
		if err != nil {
			return err
		}
	}

	return nil
}
