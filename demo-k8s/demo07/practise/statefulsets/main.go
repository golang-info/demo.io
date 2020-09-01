package statefulsets

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GreateStatefulSets(kubeClient kubernetes.Interface, sfs *appsv1.StatefulSet, namespace string) (*appsv1.StatefulSet, error) {
	statefulsets, err := GetStatefulSets(kubeClient, namespace)
	if statefulsets != nil {
		// TODO: update statefulsets
		fmt.Println("Staefulsets already exist")
		return statefulsets, err
	}

	sfs, err = kubeClient.AppsV1().StatefulSets(namespace).Create(sfs)
	if err != nil {
		// TODO: error & logs
		return nil, errors.Errorf(err, "creating statefulsets '%s'", sfs)
	}

	fmt.Printf("Statefulsets created %s / %s", sfs.Namespace, sfs.Name)
	return sfs, nil
}

func DeleteStatefulSets(kubeClient kubernetes.Interface, namespace, id string) error {
	sfs, err := GetStatefulSets(kubeClient, id)
	if sfs == nil {
		errors.Wrapf(err, "Statefulsets not exist")
	}

	deletePolicy := metav1.DeletePropagationBackground
	err = kubeClient.AppsV1().StatefulSets(namespace).Delete(id, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	fmt.Printf("Statefulsets deleted %s / %s", sfs.Namespace, sfs.Name)
	return err
}

func GetStatefulSetsList(kubeClient kubernetes.Interface, namespace string) ([]string, error) {
	var nodeList []string
	result, err := kubeClient.AppsV1().StatefulSets(namespace).List(metav1.ListOptions{
		LabelSelector: "node",
	})
	if err != nil {
		return nil, errors.Wrapf(err, "get statefulsets list '%s'", result)
	}
	for node := 0; node < len(result.Items); node++ {
		nodeList = append(nodeList, result.Items[node].Name)
	}

	fmt.Println("Success Get  StatefulSets List")
	return nodeList, nil
}

func GetStatefulSets(kubeClient kubernetes.Interface, namespace string, id string) (*appsv1.StatefulSet, error) {
	sfs, err := kubeClient.AppsV1().StatefulSets(namespace).Get(id, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	fmt.Printf("Success Get StatefulSets %s", id)
	return sfs, nil
}
