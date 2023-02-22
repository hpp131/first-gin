package admin

import "github.com/gin-gonic/gin"

type AdminController struct {
	//通过结构体的嵌套实现controller的继承
	BaseController
}

func (c AdminController) Index(ctx *gin.Context) {
	//ctx.String(200, "admin-index api success")
	c.Success(ctx)
}

func (c AdminController) User(ctx *gin.Context) {
	name, _ := ctx.Get("name")
	result, _ := name.(string)
	ctx.String(200, result)
}
