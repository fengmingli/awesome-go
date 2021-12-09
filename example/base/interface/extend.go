/**
 * @Author LFM
 * @Date 2021/12/9 3:03 下午
 * @Since V1
 */

package _interface

import "fmt"

type BInterface interface {
	Resource() BNamespaceResourceInterface
}
type BDynamicClient struct {
}

func (c *BDynamicClient) Resource() BNamespaceResourceInterface {
	return &BDynamicResourceClient{}
}

func NewBDynamicClient() *BDynamicClient {
	return &BDynamicClient{}
}

//BResourceInterface =======================
type BResourceInterface interface {
	Create()
}

type BDynamicResourceClient struct {
	namespace string
	Name      string
}

func (c *BDynamicResourceClient) Create() {
	fmt.Println(c.namespace)
}

//BNamespaceResourceInterface ====================
type BNamespaceResourceInterface interface {
	Namespace(string) BResourceInterface
	BResourceInterface
}

func (c *BDynamicResourceClient) Namespace(ns string) BResourceInterface {
	ret := *c
	ret.namespace = ns
	return &ret
}
