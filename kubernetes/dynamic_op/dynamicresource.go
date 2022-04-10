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
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DynamicResource struct {
	client.Client
}

type DynamicResourceInterface interface {
	Get(string) error
	UpdateStatus(lvm *api.Hxlvm) error
}

func NewDynamicResource(client client.Client) *DynamicResource {
	return &DynamicResource{
		client,
	}
}

func (d *DynamicResource) Get(name string) error {
	gvr := schema.GroupVersionResource{
		Group:    "lvmsimple.hexin.io",
		Version:  "v1",
		Resource: "Hxlvm",
	}
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/lifengming/.kube/config")
	newForConfig, _ := dynamic.NewForConfig(config)
	obj, err1 := newForConfig.
		Resource(gvr).
		Namespace("default").
		Get(context.Background(), name, metav1.GetOptions{}, "")
	if nil != err1 {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}

	obj.MarshalJSON()
	return nil
}

func (d *DynamicResource) UpdateStatus(lvm *api.Hxlvm) error {
	marshal, err := json.Marshal(lvm)
	if err != nil {
		return fmt.Errorf("lvm crd json marshal err:%v", marshal)
	}

	u := &unstructured.Unstructured{}
	err1 := u.UnmarshalJSON(marshal)
	if err != nil {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}

	c, err := client.New(k8sclient.GetRestClient(), client.Options{})

	if err != nil {
		fmt.Println(err)
	}

	err2 := c.Status().Update(context.Background(), u)
	if nil != err2 {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}

	return nil
}
