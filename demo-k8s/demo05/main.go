package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"os/user"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	usr, err := user.Current()
	if err != nil {
		panic(err.Error())
	}
	if home := usr.HomeDir; home != "" {
		kubeconfig = flag.String("kubeconifg",
			filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "",
			"absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	res := clientset.RESTClient().Get().AbsPath("/healthz").Do()
	if res.Error() != nil {
		panic(res.Error().Error())
	}

	r, err := res.Raw()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s", r)
	
}