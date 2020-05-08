package main

/*
	k8s-create-deployment.go
	https://gist.github.com/ffledgling/62814ccbb12fd6b9ab6fd89da7e80cd2
*/

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	kubeconfig = "/Users/liuyanchao/.kube/config"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	watch_deploy(clientset)
}

func watch_deploy(clientset *kubernetes.Clientset) {
	w, _ := clientset.AppsV1().Deployments(apiv1.NamespaceDefault).Watch(metav1.ListOptions{})
	for {
		select {
		case e := <-w.ResultChan():
			fmt.Println(e.Type, e.Object)
			deployment := e.Object.(*appsv1.Deployment)
			fmt.Println(deployment.Name)
			fmt.Println(deployment.Status.Conditions)
		}
	}
}
