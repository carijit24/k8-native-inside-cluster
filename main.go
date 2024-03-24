package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error while loading in cluster config - %s", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handle error
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error while fetching pod list - %s", err.Error())
	}
	fmt.Println("pods in default namespace are - ")
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error while fetching deployment list - %s", err.Error())
	}
	fmt.Println("")
	fmt.Println("deployments in default namespaces are - ")
	for _, deployment := range deployments.Items {
		fmt.Printf("%s", deployment.Name)
	}
}
