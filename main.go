package main

import (
	"context"
	"github.com/Yihen-Liu/go-service/log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	//go http.RunRPCService(ctx)
	//go http.MonitorBestHeight(ctx)
	//go http.ClearRedundantTxs(ctx)
	//go lightning.StartLnTestnetService(ctx)
	for {
		select {
		case <-ctx.Done():
			log.Info("main quit")
		case <-time.After(5 * time.Second):

			//log.InitLog(conf.Conf.LogLevel, conf.Conf.LogName, log.Stdout)
			log.InitLog(2, "./logs/")
		}
	}

}
