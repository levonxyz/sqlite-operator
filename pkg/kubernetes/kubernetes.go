package kubernetes

import (
	"fmt"

	"github.com/charmbracelet/log"
	k "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Kubernetes is a wrapper for the clientset
type Kubernetes struct {
	clientset *k.Clientset
}

// New creates a new client by using a kubeconfig,
// or by using an inClusterConfig if the kubeconfig is not provided.
func New(kubeconfig string) (*Kubernetes, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error building config from flags: %w", err)
	}

	clientset, err := k.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("cannot create a config: %w", err)
	}

	log.Info("clientset successfully created")
	return &Kubernetes{clientset}, nil
}
