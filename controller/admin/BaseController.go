package admin

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (c BaseController) Success(ctx *gin.Context) {
	ctx.String(200, "success")
}

func (c BaseController) Error(ctx *gin.Context) {
	ctx.String(200, "error")
}
