package main

/**
 * @Author: LFM
 * @Date: 2023/12/22 23:33
 * @Since: 1.0.0
 * @Desc: TODO
 */

type OptFunc func(opt *Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 0,
		id:      "",
		tls:     false,
	}
}

func NewServer(opts ...OptFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		Opts: o,
	}
}

func withTLS(fn *Opts) {
	fn.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opt *Opts) {
		opt.maxConn = n
	}
}

//func main() {
//	server := NewServer(withMaxConn(123))
//	fmt.Printf("%+v\n",server)
//}
