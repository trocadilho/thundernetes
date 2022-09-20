package main

import (
	"context"
	"testing"
)

func TestNotifier(t *testing.T) {
	t.Run("notifier", func(t *testing.T) {
		notifier := NewOutOfClusterKubernetesMaintenanceNotifier("aks-agentpool-33482676-vmss000001")

		err := notifier.Notify(context.TODO())

		if err != nil {
			t.Fatalf("not expected")
		}
	})
}
