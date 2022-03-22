/**
 * @Author LFM
 * @Date 2021/10/14 8:09 下午
 * @Since V1
 */

package deal_lock

import (
	"net/http"
	"net/http/pprof"
)

const (
	pprofAddr string = ":7890"
)

func StartHTTPDebug() {
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	server := &http.Server{Addr: pprofAddr, Handler: pprofHandler}
	go server.ListenAndServe()

}
