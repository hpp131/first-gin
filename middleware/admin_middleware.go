package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AdminMiddleWare(ctx *gin.Context) {
	fmt.Println("this is admin middleware")
	// 通过ctx.set方法实现middleware与controller共享数据
	ctx.Set("name", "test")
	ctx.Next()
}
