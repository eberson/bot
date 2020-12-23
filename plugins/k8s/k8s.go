package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8s struct {
	client *kubernetes.Clientset
}

func newK8S(kubeConfigPath string) (*k8s, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)

	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	return &k8s{
		client: client,
	}, nil
}
