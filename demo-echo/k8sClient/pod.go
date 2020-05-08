package k8sClient

import (
	"demo.io/demo-echo/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodsNum() int {

	pods, err := common.ClientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items))
	return len(pods.Items)
}
