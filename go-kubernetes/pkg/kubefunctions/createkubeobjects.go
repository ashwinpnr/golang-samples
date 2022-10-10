package kubefunctions

import (
	"context"
	"fmt"
	"io/ioutil"

	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

func CreateDeployment(ctx context.Context, clientset *kubernetes.Clientset) (string, error) {
	var deployment *v1.Deployment
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	appFile, err := ioutil.ReadFile("./app.yaml")
	if err != nil {
		return "", fmt.Errorf("readfile error: %s", err)
	}

	obj, groupVersionKind, err := scheme.Codecs.UniversalDeserializer().Decode(appFile, nil, nil)
	if err != nil {
		return "", fmt.Errorf("Decode error: %s", err)
	}
	switch obj.(type) {
	case *v1.Deployment:
		deployment = obj.(*v1.Deployment)
	default:
		return "", fmt.Errorf("Unrecognized type: %s\n", groupVersionKind)
	}
	result, err := deploymentsClient.Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return "", fmt.Errorf("deployment create  error: %s", err)
	}
	return result.Name, nil

}
