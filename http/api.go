package http

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */
type Response struct {
	BlockNumbers []uint64 `json:"BlockNumbers"`
	Txs          []int    `json:"Txs"`
}

func Txs(c *gin.Context) {
	var blockNumbers []uint64
	var txs []int
	latestTxs.Range(func(key, value any) bool {
		blockNumbers = append(blockNumbers, key.(uint64))
		txs = append(txs, len(value.(types.Transactions)))
		return true
	})
	c.JSON(http.StatusOK, Response{
		BlockNumbers: blockNumbers,
		Txs:          txs,
	})
}
