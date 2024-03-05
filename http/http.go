package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Yihen-Liu/go-service/config"
	"github.com/Yihen-Liu/go-service/log"
	"github.com/gin-gonic/gin"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */

func RunRPCService(ctx context.Context) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	rest := router.Group("/v1")
	rest.GET("txs", Txs)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.CarrierConf.Backend.Host, config.CarrierConf.Backend.RpcPort),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start testnet listen service err: %s\n", err)
			return
		}
		log.Info("start testnet RPC listen service successed.")
	}()

	select {
	case <-ctx.Done():
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("shutdown testnet listen service err:", err)
		}
	}
}
