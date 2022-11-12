package response

import (
	"frp-panel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, success bool, data interface{}) {
	ctx.JSON(httpStatus, gin.H{
		"success": success,
		"data":    data,
	})
}

func ResponsePage(ctx *gin.Context, httpStatus int, success bool, data interface{}, total int) {
	ctx.JSON(httpStatus, gin.H{
		"success": success,
		"data":    data,
		"total":   total,
	})
}

func LoginResponse(ctx *gin.Context, httpStatus int, loginResult model.LoginResult) {
	ctx.JSON(httpStatus, gin.H{
		"currentAuthority": loginResult.CurrentAuthority,
		"status":           loginResult.Status,
		"type":             loginResult.Type,
		"token":            loginResult.Token,
	})
}

func Success(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, true, data)
}

func PageDataSuccess(ctx *gin.Context, data interface{}, total int) {
	ResponsePage(ctx, http.StatusOK, true, data, total)
}

func Data(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func Fail(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, false, data)
}
