package main

import (
	"context"
	"github.com/Yihen-Liu/go-service/http"
	"github.com/Yihen-Liu/go-service/log"
	"os"
	"os/signal"
	"syscall"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */

func getContext() context.Context {
	signalsToCatch := []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
	interceptor := make(chan os.Signal, 1)
	signal.Notify(interceptor, signalsToCatch...)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-interceptor
		cancel()
	}()
	return ctx
}

func main() {
	ctx := getContext()

	go http.RunRPCService(ctx)
	//go lightning.StartLnTestnetService(ctx)
	select {
	case <-ctx.Done():
		log.Info("main quit")
	}

}
