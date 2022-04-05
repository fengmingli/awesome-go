package dynamic

import (
	"awesome-go/kubernetes/dynamic/lvmsimple/api"
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

/**
 * @Author: LFM
 * @Date: 2022/4/5 3:26 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

type DynamicResource struct {
	Client dynamic.Interface
}

type DynamicResourceInterface interface {
	Get()
	UpdateStatus(lvm *api.Hxlvm) error
}

func NewDynamicResource(dynamicClient dynamic.Interface) *DynamicResource {
	return &DynamicResource{
		Client: dynamicClient,
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

	gvr := schema.GroupVersionResource{
		Group:    lvm.GroupVersionKind().Group,
		Version:  lvm.GroupVersionKind().Version,
		Resource: lvm.Kind,
	}

	u := &unstructured.Unstructured{}
	err1 := u.UnmarshalJSON(marshal)
	if err != nil {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}

	_, err2 := d.Client.
		Resource(gvr).
		Namespace(lvm.GetNamespace()).
		UpdateStatus(context.Background(), u, metav1.UpdateOptions{})

	if nil != err2 {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}

	return nil
}
