package chain

import (
	"fmt"
	"testing"
)

/**
 * @Author: LFM
 * @Date: 2022/11/3 23:35
 * @Since: 1.0.0
 * @Desc: TODO
 */

func MemoryGetValue(key interface{}) (interface{}, bool) {
	return fmt.Sprintf("%s%s", key, "MemoryGetValue"), false
}

func RedisGetValue(key interface{}) (interface{}, bool) {
	return fmt.Sprintf("%s%s", key, "RedisGetValue"), false
}

func MysqlGetValue(key interface{}) (interface{}, bool) {
	return fmt.Sprintf("%s%s", key, "MysqlGetValue"), true
}

func TestChain(t *testing.T) {
	value := GetResultByChain("key", MemoryGetValue, RedisGetValue, MysqlGetValue)
	t.Log(value)
}
