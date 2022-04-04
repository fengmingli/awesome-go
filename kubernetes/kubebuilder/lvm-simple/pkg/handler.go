package pkg

import (
	"context"
	"fmt"
	v1 "hexin.io/lvm-simple/api/v1"
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

	name := "pvc-dsadsa-dsadasd-dasd33"
	if len(lv2.Status.PvItems) == 0 {
		lv2.Status.PvItems = map[string]v1.Item{}
	}

	item := v1.Item{}
	i := lv2.Status.PvItems[name]
	if i != item {
		item.Name = "pvc-dsadsa-dsadasd-dasd3eeee3"
		item.Size = 10
		item.SizeDesc = "10G"
		lv2.Status.PvItems[item.Name] = item
	}
	if err2 := h.Status().Update(context.TODO(), lv2); err2 != nil {
		fmt.Println(err2, "failed to update status", "name", crd.Name, "uid", crd.UID)
	}
}

