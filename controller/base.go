package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tmscube-go/constant"
)

type response struct {
	Code constant.ResponseCode `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
}

func setContent(ctx *gin.Context, statusCode int, resp response) {
	ctx.Set("response", resp)
	ctx.JSON(statusCode, resp)
	return
}

func Success(ctx *gin.Context, data interface{}) {
	resp := response{
		Code: constant.SUCCESS,
		Msg:  "ok",
		Data: data,
	}
	setContent(ctx, http.StatusOK, resp)
}

func Error(ctx *gin.Context, statusCode int, desc string) {
	resp := response{
		Code: constant.ERROR,
		Msg:  desc,
		Data: nil,
	}
	setContent(ctx, http.StatusOK, resp)
}
