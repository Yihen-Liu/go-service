package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Yihen-Liu/go-service/common/web3"
	"github.com/Yihen-Liu/go-service/config"
	"github.com/Yihen-Liu/go-service/db"
	"github.com/Yihen-Liu/go-service/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
var bestHeight int64
var latestTxs = new(sync.Map)

type txJSON struct {
	Type hexutil.Uint64 `json:"type"`

	ChainID              *hexutil.Big      `json:"chainId,omitempty"`
	Nonce                *hexutil.Uint64   `json:"nonce"`
	To                   *common.Address   `json:"to"`
	Gas                  *hexutil.Uint64   `json:"gas"`
	GasPrice             *hexutil.Big      `json:"gasPrice"`
	MaxPriorityFeePerGas *hexutil.Big      `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *hexutil.Big      `json:"maxFeePerGas"`
	MaxFeePerBlobGas     *hexutil.Big      `json:"maxFeePerBlobGas,omitempty"`
	Value                *hexutil.Big      `json:"value"`
	Input                *hexutil.Bytes    `json:"input"`
	AccessList           *types.AccessList `json:"accessList,omitempty"`
	BlobVersionedHashes  []common.Hash     `json:"blobVersionedHashes,omitempty"`
	V                    *hexutil.Big      `json:"v"`
	R                    *hexutil.Big      `json:"r"`
	S                    *hexutil.Big      `json:"s"`
	YParity              *hexutil.Uint64   `json:"yParity,omitempty"`

	// Only used for encoding:
	Hash common.Hash `json:"hash"`
}
type Persistment struct {
	BlockNumbers []int64                  `json:"BlockNumbers"`
	Txs          []web3.BscRpcTransaction `json:"Txs"`
}

func init() {
	//1. 从leveldb里取出内容
	db.LDB.Get([]byte("persist"), nil)
	//2. 获得level里的最新高度
	//3. 从链上拿到最新的高度
	//4. 比较上述两者的差,补齐丢失的块, 写入latestTxs中
}

func ClearRedundantTxs(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Minute):
			currentHeight, err := web3.LatestBlockNumber(config.CarrierConf.BscChain.Rpc, "bsc")
			if err != nil {
				log.Errorf("get best block number err:", err.Error())
				continue
			}

			if currentHeight != bestHeight {
				continue
			} //只在区块链高度和内存高度一致时，才继续处理

			minHeight := bestHeight
			latestTxs.Range(func(key, value any) bool {
				if minHeight > key.(int64) {
					minHeight = key.(int64)
				}
				return true
			})

			count := 0
			for i := bestHeight; i >= minHeight; i-- {
				value, ok := latestTxs.Load(i)
				if !ok { //有可能另外一个协程刚更新了bestHeight，但是还没有获取区块内容, 导致Load失效
					continue
				}
				count += len(value.([]web3.BscRpcTransaction))
				if count >= 8 {
					for j := minHeight; j < i; j++ {
						log.Infof("delete block height:%d", j)
						latestTxs.Delete(j)
					}
					break
				}
			}

			var blockNumbers []int64
			var txs []web3.BscRpcTransaction
			latestTxs.Range(func(key, value any) bool {
				if len(value.([]web3.BscRpcTransaction)) > 0 {
					blockNumbers = append(blockNumbers, key.(int64))
					txs = append(txs, value.([]web3.BscRpcTransaction)...)
				}
				return true
			})

			if len(blockNumbers) > 0 {
				var data []byte
				if data, err = json.Marshal(Persistment{BlockNumbers: blockNumbers, Txs: txs}); err != nil {
					log.Errorf("marshal persistment err:%s", err.Error())
					return
				}
				if err = db.LDB.Put([]byte("persist"), data, nil); err != nil {
					log.Errorf("leveldb put persist err:%s", err.Error())
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
func doClientRequestBlock() {
	client, err := ethclient.Dial(config.CarrierConf.BscChain.Rpc)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-time.After(1 * time.Second):
			height, err := client.BlockNumber(context.Background())
			if err != nil {
				log.Errorf("get block height err:", err.Error())
				continue
			}

			block, err := client.BlockByNumber(context.Background(), new(big.Int).SetInt64(int64(height)))
			if err != nil {
				log.Errorf("get block err:", err.Error())
				continue
			}

			txs := block.Transactions()
			if len(txs) == 0 {
				continue
			}

			for _, tx := range txs {
				innerBytes, err := tx.MarshalBinary()
				if err != nil {
					log.Errorf("marshal tx.inner err:", err.Error())
					continue
				}
				tx.UnmarshalJSON(innerBytes)
				tx.AccessList()
			}
		}
	}
}
func MonitorBestHeight(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			currentHeight, err := web3.LatestBlockNumber(config.CarrierConf.BscChain.Rpc, "bsc")
			if err != nil {
				log.Errorf("get best block number err:", err.Error())
				continue
			}
			if currentHeight == bestHeight+1 || bestHeight == 0 {
				bestHeight = currentHeight

				block, err := web3.GetBscBlockByNum(config.CarrierConf.BscChain.Rpc, fmt.Sprintf("0x%x", bestHeight))
				if err != nil {
					log.Errorf("get block err:%s, block number:%d", err.Error(), bestHeight)
				}
				log.Infof("get block, block number:%d", bestHeight)
				latestTxs.Store(bestHeight, block.Result.Transactions)
			}

			if currentHeight > bestHeight+1 { //说明中间丢块了，要把丢失的块找回来
				for i := bestHeight + 1; i <= currentHeight; i++ {
					log.Infof("get lossing block, block number:%d", i)
					block, err := web3.GetBscBlockByNum(config.CarrierConf.BscChain.Rpc, fmt.Sprintf("0x%x", i))
					if err != nil {

					}
					latestTxs.Store(i, block.Result.Transactions)
				}
			}

		case <-ctx.Done():
			return
		}
	}
}
