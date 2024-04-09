/**
 * @Author LFM
 * @Date 2021/10/14 8:52 下午
 * @Since V1
 */

package analysis

import "testing"

func TestEscapeAnalysisExample(t *testing.T) {
	manager := NewWorkerManager()
	factory := manager.DynamicDataSourceFactory("A")
	factory.Split()

	factoryB := manager.DynamicDataSourceFactory("B")
	factoryB.Send()

	factoryA2 := manager.DynamicDataSourceFactory("A")
	factoryA2.Send()

	factoryB2 := manager.DynamicDataSourceFactory("B")
	factoryB2.Split()
}
