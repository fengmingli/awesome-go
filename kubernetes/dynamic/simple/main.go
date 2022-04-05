package main

import (
	"awesome-go/kubernetes/dynamic"
	c "k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

/**
 * @Author: LFM
 * @Date: 2022/4/5 3:57 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/lifengming/.kube/config")
	newForConfig, _ := c.NewForConfig(config)
	dynamicResource := dynamic.NewDynamicResource(newForConfig)

	dynamicResource.Get("hxlvm-sample-1")
}
