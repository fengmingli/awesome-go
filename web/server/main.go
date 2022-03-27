package main

import (
	logger "awesome-go/tools/log"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

/**
 * @Author: LFM
 * @Date: 2022/3/27 7:49 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

func main() {
	logger.InitZeroLog()
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		err1 := json.NewEncoder(w).Encode("123456")
		if err1 != nil {
			return
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Info().Msgf("http lister 8080 start fail,errMsg:%v", err)
	}
}
