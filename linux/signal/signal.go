package main

/**
 * @Author: LFM
 * @Date: 2022/4/16 11:23 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	cmd := exec.Command("ping")

	cmd.Args = []string{"ping", "baidu.com"}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) //handle signal during grace period
	go func() {
		sig := <-signalChan
		fmt.Printf("catch kubernetes signal:%v, stop cmd:%v \n", sig, cmd.Args)
		err := cmd.Process.Signal(os.Interrupt)
		if err != nil {
			return
		}
	}()

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		os.Exit(1)
	}
}
