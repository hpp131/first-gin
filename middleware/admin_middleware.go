package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MiddleWareOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is first middleware")
	}
}

func MiddleWareTwo(ctx *gin.Context) {
	fmt.Println("this is second middleware")
	// 通过ctx.set方法实现middleware与controller共享数据
	ctx.Next()
	fmt.Println("second middleware end")
}
