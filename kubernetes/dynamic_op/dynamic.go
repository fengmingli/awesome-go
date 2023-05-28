package dynamic_op

import (
	"awesome-go/kubernetes/dynamic_op/lvmsimple/api"
	"awesome-go/kubernetes/k8sclient"
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	// Get a config to talk to the apiserver
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

func updateLvmStatus(lvm *api.Hxlvm) error {
	marshal, err := json.Marshal(lvm)
	if err != nil {
		return fmt.Errorf("lvm crd json marshal err:%v", marshal)
	}
	gvr := schema.GroupVersionResource{
		Group:    lvm.GroupVersionKind().Group,
		Version:  lvm.GroupVersionKind().Version,
		Resource: "hxlvms",
	}

	u := &unstructured.Unstructured{}
	err1 := u.UnmarshalJSON(marshal)
	if err != nil {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}
	ops := metav1.UpdateOptions{}

	_, err2 := k8sclient.GetDynamicClient().
		Resource(gvr).
		Namespace("default").
		UpdateStatus(context.Background(), u, ops)

	if nil != err2 {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}
	return nil

}
