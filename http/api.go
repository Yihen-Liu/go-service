package http

import (
	"github.com/Yihen-Liu/go-service/common/web3"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */
type Response struct {
	BlockNumbers []int64                  `json:"BlockNumbers"`
	Cnts         []int                    `json:"Cnts"`
	Txs          []web3.BscRpcTransaction `json:"Txs"`
}

func Txs(c *gin.Context) {
	var blockNumbers []int64
	var txs []web3.BscRpcTransaction
	var cnts []int
	latestTxs.Range(func(key, value any) bool {
		blockNumbers = append(blockNumbers, key.(int64))
		cnts = append(cnts, len(value.([]web3.BscRpcTransaction)))
		if len(value.([]web3.BscRpcTransaction)) > 0 {
			txs = append(txs, value.([]web3.BscRpcTransaction)...)
		}
		return true
	})
	c.JSON(http.StatusOK, Response{
		BlockNumbers: blockNumbers,
		Cnts:         cnts,
		Txs:          txs,
	})
}
