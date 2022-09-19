package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type MaintenanceNotifier interface {
	Notify() error
}

type KubernetesMaintenanceNotifier struct{
	clientset *kubernetes.Clientset
}

func NewKubernetesMaintenanceNotifier() MaintenanceNotifier {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	k := KubernetesMaintenanceNotifier{}
	k.clientset = clientset
	return k
}

func (n KubernetesMaintenanceNotifier) Notify() error {
	nodeName, ok := os.LookupEnv("MY_NODE_NAME")
	if !ok {
		return fmt.Errorf("MY_NODE_NAME is not present")
	} else {
		payload := []patchStringValue{{
			Op:    "replace",
			Path:  "/spec/unschedulable",
			Value: true,
		}}
		payloadBytes, _ := json.Marshal(payload)
		_, err := n.clientset.CoreV1().Nodes().Patch(context.TODO(), nodeName, types.JSONPatchType, payloadBytes, metav1.PatchOptions{})
		return err
	}
}
