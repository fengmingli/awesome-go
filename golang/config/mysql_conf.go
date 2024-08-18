package config

import "fmt"

/**
 * @Author: LFM
 * @Date: 2024/8/18 23:02
 * @Since: 1.0.0
 * @Desc: TODO
 */

type server struct {
	addr string //必须
	port int

	timeout int //可选
	maxConn int
}

type option func(s *server)

func SetTimeout(timeout int) option {
	return func(s *server) {
		s.timeout = timeout
	}
}

func SetMaxConn(max int) option {
	return func(s *server) {
		s.maxConn = max
	}
}

func NewServer(addr string, port int, options ...option) *server {
	s := &server{
		addr: addr,
		port: port,
	}

	for _, op := range options {
		op(s)
	}

	return s
}

func (s *server) Start() {
	fmt.Println("Running...")
}
