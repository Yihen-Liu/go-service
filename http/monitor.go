package http

import (
	"context"
	"github.com/Yihen-Liu/go-service/config"
	"github.com/Yihen-Liu/go-service/log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
	"time"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */
var bestHeight uint64
var latestTxs = new(sync.Map)

func ClearRedundantTxs(ctx context.Context) {
	client, err := ethclient.Dial(config.CarrierConf.BscChain.Rpc)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case <-time.After(1 * time.Minute):
			currentHeight, err := client.BlockNumber(context.Background())
			if err != nil {
				log.Errorf("get best block number err:", err.Error())
				continue
			}

			if currentHeight != bestHeight {
				continue
			} //只在区块链高度和内存高度一致时，才继续处理

			var keys []uint64
			minHeight := bestHeight
			latestTxs.Range(func(key, value any) bool {
				keys = append(keys, key.(uint64))
				if minHeight > key.(uint64) {
					minHeight = key.(uint64)
				}
				return true
			})

			count := 0
			for i := bestHeight; i >= minHeight; i-- {
				value, _ := latestTxs.Load(i)
				count += len(value.(types.Transactions))
				if count >= 8 {
					for j := minHeight; j < i; j++ {
						latestTxs.Delete(j)
					}
					break
				}
			}

		case <-ctx.Done():
			return
		}
	}
}

func MonitorBestHeight(ctx context.Context) {
	client, err := ethclient.Dial(config.CarrierConf.BscChain.Rpc)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-time.After(1 * time.Second):
			currentHeight, err := client.BlockNumber(context.Background())
			if err != nil {
				log.Errorf("get best block number err:", err.Error())
				continue
			}
			if currentHeight == bestHeight+1 || bestHeight == 0 {
				bestHeight = currentHeight

				block, err := client.BlockByNumber(context.Background(), new(big.Int).SetInt64(int64(bestHeight)))
				if err != nil {
					log.Errorf("get block err:%s, block number:%d", err.Error(), bestHeight)
				}
				log.Infof("get block, block number:%d", bestHeight)
				latestTxs.Store(bestHeight, block.Transactions())
			}

			if currentHeight > bestHeight+1 { //说明中间丢块了，要把丢失的块找回来
				for i := bestHeight + 1; i <= currentHeight; i++ {
					log.Infof("get lossing block, block number:%d", i)
					block, err := client.BlockByNumber(context.Background(), new(big.Int).SetInt64(int64(i)))
					if err != nil {

					}
					latestTxs.Store(i, block.Transactions())
				}
			}

		case <-ctx.Done():
			return
		}
	}
}
