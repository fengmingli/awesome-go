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
)

func main() {
	//
	//resource := "hxlvms.lvmsimple.hexin.io"
	//// 同样这里也是 DynamicSharedInformerFactory, 而不是 SharedInformerFactory
	//informerFactory := dynamicinformer.NewDynamicSharedInformerFactory(
	//	k8sclient.GetDynamicClient(), 0)
	//
	//// 通过 schema 包提供的 ParseResourceArg 由资源描述字符串解析出 GroupVersionResource
	//gvr, _ := schema.ParseResourceArg(resource)
	//if gvr == nil {
	//	fmt.Println("cannot parse gvr")
	//	return
	//}
	//unstructuredList, _err := k8sclient.GetDynamicClient().
	//	Resource(*gvr).
	//	Namespace("default").
	//	List(context.TODO(), metav1.ListOptions{Limit: 100})
	//if _err != nil {
	//	fmt.Println(_err)
	//}
	//fmt.Println(unstructuredList)
	//// 使用 gvr 动态生成 Informer
	//informer := informerFactory.ForResource(*gvr).Informer()
	//// 熟悉的代码，熟悉的味道，只是收到的 obj 类型好像不太一样
	//informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc: func(obj interface{}) {
	//		// *unstructured.Unstructured 类是所有 Kubernetes 资源类型公共方法的抽象，
	//		// 提供所有对公共属性的访问方法，像 GetName, GetNamespace, GetLabels 等等，
	//		s, ok := obj.(*api.Hxlvm)
	//		if !ok {
	//			return
	//		}
	//		fmt.Printf("created: %s\n", s.GetName())
	//	},
	//	UpdateFunc: func(oldObj, newObj interface{}) {
	//		oldS, ok1 := oldObj.(*unstructured.Unstructured)
	//		newS, ok2 := newObj.(*unstructured.Unstructured)
	//		if !ok1 || !ok2 {
	//			return
	//		}
	//		// 要访问公共属性外的字段，可以借助 unstructured 包提供的一些助手方法：
	//		oldColor, ok1, err1 := unstructured.NestedString(oldS.Object, "spec", "color")
	//		newColor, ok2, err2 := unstructured.NestedString(newS.Object, "spec", "color")
	//		if !ok1 || !ok2 || err1 != nil || err2 != nil {
	//			fmt.Printf("updated: %s\n", newS.GetName())
	//		}
	//		fmt.Printf("updated: %s, old color: %s, new color: %s\n", newS.GetName(), oldColor, newColor)
	//	},
	//	DeleteFunc: func(obj interface{}) {
	//		s, ok := obj.(*unstructured.Unstructured)
	//		if !ok {
	//			return
	//		}
	//		fmt.Printf("deleted: %s\n", s.GetName())
	//	},
	//})
	//
	//stopCh := make(chan struct{})
	//defer close(stopCh)
	//
	//fmt.Println("Start syncing....")
	//
	//go informerFactory.Start(stopCh)
	//
	//<-stopCh
	//updateLvmStatus()
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