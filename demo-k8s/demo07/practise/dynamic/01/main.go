package main

/*
	https://erwinvaneyk.nl/kubernetes-unstructured-to-typed/
*/

import (
	"flag"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/cluster-api/api/v1alpha2"
)

func main() {
	// Create a new dynamic client
	//restConfig, err := clientcmd.BuildConfigFromFlags("", "")
	//assertNoError(err)
	//TODO: panic: /private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_demo_io_demo_k8s_demo07_practise_dynamic_01 flag redefined: kubeconfig
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	restConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	kubeClient, err := dynamic.NewForConfig(restConfig)
	assertNoError(err)

	// Get a resource (returns an unstructured object)
	resourceScheme := v1alpha2.SchemeBuilder.GroupVersion.WithResource("cluster")
	resp, err := kubeClient.Resource(resourceScheme).Namespace("All").Get("app", metav1.GetOptions{})
	assertNoError(err)

	// Convert the unstructured object to cluster.
	unstructured := resp.UnstructuredContent()
	var cluster v1alpha2.Cluster
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &cluster)
	assertNoError(err)

	// Use the typed object.
	fmt.Println(cluster.Status.Phase)


}

func assertNoError(err error) {
	if err != nil {
		panic(err)
	}
}