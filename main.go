package main

import (
	"first-gin/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
	"time"
)

type Result struct {
	Name string
	Num  int
}

type Bind struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 设置中间件，最后一个func方法前触发的方法都可以叫做中间件，也就是说中间件可以定义并使用多个
func middleFunc(ctx *gin.Context) {
	start_time := time.Now().Unix()
	ctx.String(200, "I'm a middleware")
	ctx.Next()
	end_time := time.Now().Unix()
	ctx.String(200, strconv.Itoa(int(end_time-start_time)))
}

func main() {
	r := gin.Default()
	//引入抽离出来的路由分组
	routes.AdminRouterGroup(r)
	// 注册自定义模版函数
	r.SetFuncMap(template.FuncMap{
		"formatDate": CustomTemplateFunc,
	})
	//使用HTML渲染时需要使用以下配置
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	user := Result{
		Name: "test_user",
		Num:  100,
	}
	r.GET("/", func(context *gin.Context) {
		context.String(200, "value is %v", "你好gin")
	})

	r.GET("/json", func(context *gin.Context) {
		// gin.H 等效于map[string]interface{};此处也可传入结构体
		context.JSON(200, gin.H{"msg": "ok"})
	})
	r.GET("/json2", func(context *gin.Context) {
		context.JSON(200, Result{Name: "gin-json", Num: 1})
	})
	r.GET("/html", func(context *gin.Context) {
		// 使用html模版渲染
		context.HTML(200, "admin/index.html", gin.H{
			"title": "this is index.html page",
		})
	})
	r.GET("/default", func(context *gin.Context) {
		// 使用html模版渲染
		context.HTML(200, "default/index.html", gin.H{
			// 可以支持传入结构体/结构体中的某字段
			"title": user.Name,
		})
	})
	r.GET("/customFunc", func(context *gin.Context) {
		// 使用html模版渲染
		context.HTML(200, "default/index.html", gin.H{
			// 可以支持传入结构体/结构体中的某字段
			"title": time.Now(),
		})
	})
	r.GET("/getvalue", func(context *gin.Context) {
		uid := context.Query("uid")
		context.String(200, fmt.Sprintf("uid is %v", uid))
	})
	//获取GET POST form表单传值
	r.GET("/postvalue", func(context *gin.Context) {
		context.HTML(200, "default/add_user.html", gin.H{})
	})
	r.POST("/doAddUser", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(200, gin.H{
			"username":  username,
			"password:": password,
		})
	})
	// 使用ShouldBind将请求入参绑定到结构体
	r.GET("/bindvalue", func(context *gin.Context) {
		var bind Bind
		context.ShouldBind(&bind)
		context.JSON(200, bind)
	})
	r.GET("/middleware", middleFunc, func(context *gin.Context) {
		time.Sleep(time.Second)
		context.JSON(200, gin.H{
			"msg": "调用middleware",
		})
	})

	r.Run()
}

func CustomTemplateFunc(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", year, month, day)
}
