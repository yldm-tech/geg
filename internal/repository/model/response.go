package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResSuccess(c *gin.Context, data interface{}) {
	result := Response{
		Code: http.StatusOK,
		Msg:  "SUCCESS",
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}

// ResFailure ResFailure
func ResFailure(c *gin.Context, code uint32, message string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  message,
		Data: nil,
	})
}

// ResFailureV2 only used in the API for content provider
func ResFailureV2(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code: uint32(code),
		Msg:  message,
		Data: nil,
	})
}
