package main

import (
	"github.com/gin-gonic/gin"
)

type Result struct {
	Name string
	Num  int
}

func main() {
	r := gin.Default() //
	r.GET("/", func(context *gin.Context) {
		context.String(200, "value is %v", "你好gin")
	})
	r.GET("/news", func(context *gin.Context) {
		context.String(200, "value is %v", "新闻1111")
	})
	r.POST("/create", func(context *gin.Context) {
		context.String(200, "value is %v", "create新闻")
	})
	r.GET("/json", func(context *gin.Context) {
		// gin.H 等效于map[string]interface{};此处也可传入结构体
		context.JSON(200, gin.H{"msg": "ok"})
	})
	r.GET("/json2", func(context *gin.Context) {
		// gin.H 等效于map[string]interface{};此处也可传入结构体
		context.JSON(200, Result{Name: "gin-json", Num: 1})
	})
	r.Run()
}
