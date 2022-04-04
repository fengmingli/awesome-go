package k8sclient

import (
	"flag"
	"fmt"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var client *kubernetes.Clientset
var dynamicClient dynamic.Interface

func init() {
	var kubeconfig *string
	// 获取k8s配置文件kubeconfig的地址
	if home := homeDir(); home != "" { // 接受命令行中的name参数, go run main.go --kubeconfig=xxx
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	fmt.Println(*kubeconfig)

	// 使用k8s.io/k8sclient-go/tools/clientcmd生成config的对象
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	/// 使用k8s.io/k8sclient-go/kubernetes生成一个ClientSet的客户端
	// 客户端生成后，就可以使用这个客户端与k8s API server进行交互，如Create/Retrieve/Update/Delete Resource
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	client = clientset
	forConfig, err := dynamic.NewForConfig(config)
	if err != nil {

	}
	dynamicClient = forConfig

}

func GetClientSet() *kubernetes.Clientset {
	return client
}

func GetDynamicClient() dynamic.Interface {
	return dynamicClient
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	return os.Getenv("USERPROFILE") // windows
}
