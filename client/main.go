package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// restClient
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	//if err != nil {
	//	panic(err)
	//}
	//config.GroupVersion = &v1.SchemeGroupVersion
	//config.NegotiatedSerializer = scheme.Codecs
	//config.APIPath = "/api"
	//restClient, err := rest.RESTClientFor(config)
	//if err != nil {
	//	panic(err)
	//}
	//pod := v1.Pod{}
	//err = restClient.Get().Namespace("default").Resource("pods").Name("my-chart-repo-chartmuseum-7cbb59d995-8k9np").Do(context.TODO()).Into(&pod)
	//if err != nil {
	//	panic(err)
	//} else {
	//	println(pod.Name)
	//}

	//	clientSet
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	//if err != nil {
	//	panic(err)
	//}
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	//coreV1 := clientset.CoreV1()
	//podList, err := coreV1.Pods("kube-system").List(context.TODO(), v1.ListOptions{})
	//if err != nil {
	//	panic(err)
	//} else {
	//	for i, p := range podList.Items {
	//		println(i, p.Name)
	//	}
	//}

	// dynamic clientset
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	resource := client.Resource(schema.GroupVersionResource{Group: "kubeeye.kubesphere.io", Version: "v1alpha1", Resource: "clusterinsights"})
	unstructured, err := resource.Get(context.TODO(), "clusterinsight-sample", v1.GetOptions{})
	if err != nil {
		panic(err)
	} else {
		obj, err := resource.Get(context.TODO(), unstructured.GetName(), v1.GetOptions{})
		if err != nil {
			panic(err)
		}
		for k, v := range obj.UnstructuredContent() {
			fmt.Printf("%v, %v\n", k, v)
		}
		//runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent())

	}
}
