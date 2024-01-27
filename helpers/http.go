package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func Success(result any, ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"result": result,
		"error":  nil,
	})
}

func Throw(err error, ctx *gin.Context, fault Fault) {
	ctx.JSON(fault.Int(), gin.H{
		"result": nil,
		"error":  err.Error(),
	})
}

func UnimplementedStub(ctx *gin.Context) {
	Throw(errors.New("unimplemented API feature"), ctx, ServerFault)
}
