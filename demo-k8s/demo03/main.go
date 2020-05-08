package main

/*
	k8s-create-deployment.go
	https://gist.github.com/ffledgling/62814ccbb12fd6b9ab6fd89da7e80cd2
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	// We use pretty instead of the common go-spew or pretty-printing because,
	// go-spew is vendored in client-go and causes problems

	"github.com/kr/pretty"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//v1beta "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

var (
	kubeconfig = "/Users/liuyanchao/.kube/config"
	dbjson     = "/Users/liuyanchao/go/src/demo.io/demo-k8s/demo01/deployment.json"
)

func main() {
	//Read the deployment json
	file, err := os.Open(dbjson)
	if err != nil {
		panic(err.Error())
	}
	dec := json.NewDecoder(file)

	//Parse it into the internal k8s structs
	var dep appsv1.Deployment
	dec.Decode(&dep)

	//Dump the struct in case you want to see what it looks like
	pretty.Println(dep)

	//dep.Namespace = "kube-public"

	//pretty.Println(dep)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//deploymentclient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)

	// List existing deployments in namespace
	deployments, err := deploymentclient.List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	} else {
		for i, e := range deployments.Items {
			log.Printf("Deployment #%d\n", i)
			log.Printf("%s", e.ObjectMeta.SelfLink)
		}
	}

	// Create a new deployment based on our template.
	//if result, err := deploymentsClient.Create(&dep); err != nil {
	//	log.Println("ERROR:", err.Error())
	//} else {
	//	pretty.Println(result)
	//}

	// https://www.e-learn.cn/content/wangluowenzhang/2190180
	watch_deploy(clientset)
}

//func watch_deploy(deploymentclient v1beta.DeploymentInterface) {
//	w, _ := deploymentclient.Watch(metav1.ListOptions{})
//	for {
//		select {
//			case e := <- w.ResultChan():
//				fmt.Println(e.Type, e.Object)
//				fmt.Printf("%T\n", e.Object)
//				fmt.Sprintf("%v+", e.Object)
//		}
//	}
//}

func watch_deploy(clientset *kubernetes.Clientset) {
	w, _ := clientset.AppsV1().Deployments(apiv1.NamespaceDefault).Watch(metav1.ListOptions{})
	for {
		select {
		case e := <- w.ResultChan():
			fmt.Println(e.Type, e.Object)
			deployment := e.Object.(*appsv1.Deployment)
			fmt.Println(deployment.Name)
			fmt.Println(deployment.Status.Conditions)
		}
	}
}