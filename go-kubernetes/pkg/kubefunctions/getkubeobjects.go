package kubefunctions

import (
	"context"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetKubeClient() (*kubernetes.Clientset, error) {

	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, err
}

func GetNameSpaceList(ctx context.Context, clientset *kubernetes.Clientset) ([]string, error) {

	var namespaces []string
	namespaceList, err := clientset.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, namespace := range namespaceList.Items {
		namespaces = append(namespaces, namespace.Name)
	}
	return namespaces, nil

}

func GetPods(ctx context.Context, clientset *kubernetes.Clientset, namespaces []string) (map[string][]string, error) {
	pods_details := make(map[string][]string)
	for _, namespace := range namespaces {
		var listofPods []string
		podList, err := clientset.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for _, pods := range podList.Items {
			listofPods = append(listofPods, pods.Name)
		}
		_, ok := pods_details[namespace]

		if ok {
			pods_details[namespace] = append(pods_details[namespace], listofPods...)
		} else {
			pods_details[namespace] = listofPods
		}

	}
	return pods_details, nil
}
