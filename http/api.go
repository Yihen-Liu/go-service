package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */
type Response struct {
	Amount  int64  `json:"amount"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Informations(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Amount:  1024,
		Code:    0,
		Message: "Success",
	})
}
