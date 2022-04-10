package main

import (
	"awesome-go/kubernetes/dynamic_op"
	"awesome-go/kubernetes/k8sclient"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/**
 * @Author: LFM
 * @Date: 2022/4/5 3:57 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

func main() {
	c, err := client.New(k8sclient.GetRestClient(), client.Options{})
	if err != nil {
	}

	dynamicResource := dynamic_op.NewDynamicResource(c)

	_ = dynamicResource.Get("hxlvm-sample-1")
	//dynamicResource.UpdateStatus(lvm)
}
