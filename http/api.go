package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */
func Txs(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}
