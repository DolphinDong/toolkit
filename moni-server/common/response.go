package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []interface{} `json:"data"`
}

// 返回500
func ResponseInternalServerErrorWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:  msg,
	})
}

// 返回200并携带msg
func ResponseOkWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  msg,
	})
}
func ResponseOKWithCodeMsg(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
}

// 返回403
func ResponseForbidden(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusForbidden, Response{
		Code: http.StatusForbidden,
		Msg:  msg,
	})
}
func ResponseTableErrorWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, LayuiTable{
		Code: LayuiTabelDataError,
		Msg:  msg,
	})
}
func ResponseInvalidParameter(ctx *gin.Context) {
	ResponseOKWithCodeMsg(ctx, http.StatusBadRequest, "invalid parameter")
}

func ResponseAddSuccess(ctx *gin.Context) {
	ResponseOkWithMsg(ctx, "添加成功")
}
func ResponseUpdateSuccess(ctx *gin.Context) {
	ResponseOkWithMsg(ctx, "修改成功")
}
func ResponseDeleteSuccess(ctx *gin.Context) {
	ResponseOkWithMsg(ctx, "删除成功")
}
