package util

import (
	"GCloud/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ctx *gin.Context
}

type HttpResult struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
}

func NewResponse(c *gin.Context)*Response  {
	return &Response{c}
}
func (r *Response) ToResponse(code int,data interface{})  {
	r.ctx.JSON(http.StatusOK,HttpResult{
		Code:code,
		Data:data,
		Msg:errcode.GetErrMessage(code),
	})
}
func (r *Response) FileToResponse(code int,data interface{})  {
	r.ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"chunkList":data,
		"msg":errcode.GetErrMessage(code),
	})
}
