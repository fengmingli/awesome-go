/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"hexin.io/lvm-simple/pkg"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	lvmsimplev1 "hexin.io/lvm-simple/api/v1"
)

// HxlvmReconciler reconciles a Hxlvm object
type HxlvmReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=lvmsimple.hexin.io,resources=hxlvms,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lvmsimple.hexin.io,resources=hxlvms/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lvmsimple.hexin.io,resources=hxlvms/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Hxlvm object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *HxlvmReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	crd := new(lvmsimplev1.Hxlvm)
	if err := r.Get(ctx, req.NamespacedName, crd); err != nil {
		if !apierrs.IsNotFound(err) {
			return ctrl.Result{}, err
		}
	}
	pkg.NewHandler(r.Client).Do(crd)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HxlvmReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&lvmsimplev1.Hxlvm{}).
		Complete(r)
}
