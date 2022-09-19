package main

type MaintenanceNotifier interface {
	Notify() error
}

type KubernetesMaintenanceNotifier struct{}

func NewKubernetesMaintenanceNotifier() MaintenanceNotifier {
	return KubernetesMaintenanceNotifier{}
}

func (n KubernetesMaintenanceNotifier) Notify() error {
	// TODO: Call Kubernetes APIs
	return nil
}
