package mysql

import (
	"testing"

	"awesome-go/golang/config"
)

/**
 * @Author: LFM
 * @Date: 2024/8/18 23:09
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestNewServer(t *testing.T) {

	newServer := config.NewServer("127.0.0.1", 8083,
		config.SetTimeout(100))

	newServer.Start()
}
