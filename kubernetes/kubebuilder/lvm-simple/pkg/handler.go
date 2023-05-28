package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "hexin.io/lvm-simple/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/**
 * @Author: LFM
 * @Date: 2022/4/3 9:24 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

type Handler struct {
	client.Client
}

func NewHandler(client2 client.Client) *Handler {
	return &Handler{
		Client: client2,
	}
}

type HandlerFuc interface {
	Do(crd *v1.Hxlvm)
}

func (h *Handler) Do(crd *v1.Hxlvm) {
	h.updateStatus(crd)
}

func (h *Handler) updateStatus(crd *v1.Hxlvm) {

	lv2 := crd.DeepCopy()

	if len(lv2.Status.PvItems) == 0 {
		lv2.Status.PvItems = map[string]v1.Item{}
	}
	if len(lv2.Status.PvcItems) == 0 {
		lv2.Status.PvcItems = map[string]v1.Item{}
	}
	name := "pvc-xxx-xxxx2"
	item := v1.Item{}
	i := lv2.Status.PvItems[name]
	if i == item {
		item.Name = name
		item.Size = 1000
		item.SizeDesc = "1000G"
		lv2.Status.PvItems[item.Name] = item
		lv2.Status.PvcItems[item.Name] = item
	}
	deleteName := "pvc-ddxzdddddd-10002"
	items := lv2.Status.PvItems
	if len(items) > 0 {
		for k := range items {
			if k == deleteName {
				delete(items, deleteName)
			}
		}
	}
	_ = updateLvmStatus(lv2)
}

func getLvmResource(lvm *v1.Hxlvm) error {
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/lifengming/.kube/config")
	newForConfig, _ := dynamic.NewForConfig(config)
	gvr := schema.GroupVersionResource{
		Group:    lvm.GroupVersionKind().Group,
		Version:  lvm.GroupVersionKind().Version,
		Resource: lvm.Kind,
	}
	obj, err1 := newForConfig.
		Resource(gvr).
		Namespace(lvm.GetNamespace()).
		Get(context.Background(), lvm.Name, metav1.GetOptions{})

	if nil != err1 {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}
	fmt.Println(obj)
	return nil
}

func updateLvmStatus(lvm *v1.Hxlvm) error {
	marshal, err := json.Marshal(lvm)
	if err != nil {
		return fmt.Errorf("lvm crd json marshal err:%v", marshal)
	}

	u := &unstructured.Unstructured{}
	err1 := u.UnmarshalJSON(marshal)
	if err != nil {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}

	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/lifengming/.kube/config")

	c, err1 := client.New(config, client.Options{})

	if err1 != nil {
		return fmt.Errorf("statusClient create fail err:%v", err1)
	}

	err2 := c.Status().Update(context.Background(), u)

	if nil != err2 {
		return fmt.Errorf("lvm crd unstructured json UnmarshalJSON err:%v", err1)
	}

	return nil
}
