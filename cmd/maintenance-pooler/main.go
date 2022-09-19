package main

import "context"

func main() {
	context := context.Background()

	metadataClient := NewAzureMetadataClient()
	notifier := NewKubernetesMaintenanceNotifier()
	checker := NewChecker(metadataClient, notifier)

	checker.Start(context)
}
