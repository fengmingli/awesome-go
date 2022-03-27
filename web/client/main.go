package main

import (
	logger "awesome-go/tools/log"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2022/3/27 7:47 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

func main() {
	logger.InitZeroLog()

	ticker := time.NewTicker(3 * time.Second)
S:
	for {
		select {
		case <-ticker.C:
			test()
			break S
		}
	}
	log.Info().Msgf("finish ")

}

func test() string {
	resp, err := http.Get("http://127.0.0.1:8080/status")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Info().Msgf("http lister 8080 start fail,errMsg:%v", err)
	}

	log.Info().Msgf("http lister 8080 start fail,errMsg:%v", string(body))
	return string(body)

}
