package main

import (
	"awesome-go/kubernetes/client"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"time"
)

func podNumOfSpecifyNode() {
	indexByPodNodeName := func(obj interface{}) ([]string, error) {
		pod, ok := obj.(*v1.Pod)
		if !ok {
			return []string{}, nil
		}
		if len(pod.Spec.NodeName) == 0 ||
			pod.Status.Phase == v1.PodSucceeded ||
			pod.Status.Phase == v1.PodFailed {
			return []string{}, nil
		}
		return []string{pod.Spec.NodeName}, nil
	}

	sharedInformers := informers.NewSharedInformerFactory(client.GetClientSet(), 0)
	podInformer := sharedInformers.Core().V1().Pods().Informer()
	err := podInformer.GetIndexer().AddIndexers(cache.Indexers{"nodeName": indexByPodNodeName})
	if err != nil {
		return
	}
	stopChan := make(chan struct{})
	defer close(stopChan)
	go podInformer.Run(stopChan)
	for range time.Tick(time.Second * 2) {
		podInformer.GetIndexer().ListKeys()
		nodeName := "docker-desktop"
		podList, err := podInformer.GetIndexer().ByIndex("nodeName", nodeName)
		if err != nil {
			klog.V(0).Info(err)

		}
		klog.V(0).Info(podList)
	}

}

func main() {
	podNumOfSpecifyNode()
}
