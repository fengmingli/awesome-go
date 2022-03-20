package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	sigs "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type NodeReconciler struct {
	client.Client
	SkipNodeFinalize bool
}

func (r *NodeReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	klog.V(0).Info("=====Reconcile===")
	klog.V(0).Info(request.Name)
	klog.V(0).Info("=================")
	klog.V(0).Info(ctx)
	return sigs.Result{}, nil
}

func main() {
	subMain()

}
func subMain() error {

	cfg, err := sigs.GetConfig()
	if err != nil {
		return err
	}
	mgr, err := sigs.NewManager(cfg, sigs.Options{})

	if err != nil {
		return err
	}

	pvccontroller := &NodeReconciler{
		Client: mgr.GetClient(),
	}

	err1 := pvccontroller.setupWithManager(mgr)
	if err1 != nil {
		return err1
	}
	// 初始化磁盘管理服务
	stopChan := make(chan struct{})
	defer close(stopChan)

	if err := mgr.Start(sigs.SetupSignalHandler()); err != nil {
		klog.Error(err, "problem running manager")
		close(stopChan)
		return err
	}
	return nil
}

func (r *NodeReconciler) setupWithManager(mgr sigs.Manager) error {

	pred := predicate.Funcs{
		CreateFunc:  func(event.CreateEvent) bool { return true },
		DeleteFunc:  func(event.DeleteEvent) bool { return false },
		UpdateFunc:  func(event.UpdateEvent) bool { return true },
		GenericFunc: func(event.GenericEvent) bool { return false },
	}

	return sigs.NewControllerManagedBy(mgr).
		WithEventFilter(pred).
		For(&corev1.Pod{}).
		Complete(r)

}
