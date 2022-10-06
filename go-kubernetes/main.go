package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ashwinpnr/golang-samples/go-kubernetes/pkg/kubefunctions"
)

func main() {
	ctx := context.Background()
	clientset, err := kubefunctions.GetKubeClient()

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	listOfNamespaces, err := kubefunctions.GetNameSpaceList(ctx, clientset)

	if err != nil {
		log.Fatalf("Error in List of namespaces : %v", err)
	}

	fmt.Printf("List of Namespace : %s \n", listOfNamespaces)

	listOfPods, err := kubefunctions.GetPods(ctx, clientset, listOfNamespaces)
	if err != nil {
		log.Fatalf("Error in list of Pods: %v", err)
	}

	for namespace, pods := range listOfPods {
		fmt.Printf("List of Pods in Namespace %s : %s \n", namespace, pods)
	}

}
