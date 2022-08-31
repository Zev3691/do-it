package service

import (
	"context"
	"os"
	"os/signal"
	"re_new/repository/mongo"
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
		defer func() {
			// 关闭mongo链接
			mongo.NewMongoDB().Disconnect(context.TODO())
		}()
		done <- struct{}{}
	}()
	return done
}
