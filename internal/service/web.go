package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	go HTTP()
	go RPC()
	<-do()
}

func do() chan struct{} {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("dddddddddd")
		done <- struct{}{}
	}()
	return done
}
