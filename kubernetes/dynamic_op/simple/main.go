package main

import (
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

/**
 * @Author: LFM
 * @Date: 2022/4/5 3:57 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

func main() {
	//c, err := client.New(k8sclient.GetRestClient(), client.Options{})
	//if err != nil {
	//}
	//
	//dynamicResource := dynamic_op.NewDynamicResource(c)
	//
	//_ = dynamicResource.Get("hxlvm-sample-1")
	//dynamicResource.UpdateStatus(lvm)
	cfg, err := config.GetConfig()
	if err != nil {
	}

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{
		Namespace: "default",
	})
	if err != nil {
	}
	fmt.Println(mgr)

}
