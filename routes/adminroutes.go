package routes

import (
	"first-gin/controller/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AdminRouterGroup(r *gin.Engine) {
	// 在路由组中配置中间件
	adminrouter := r.Group("/test")
	{
		adminrouter.GET("/index", func(context *gin.Context) {
			context.String(200, "test-index api success")
		})
		adminrouter.GET("/date", admin.AdminController{}.Date)
		adminrouter.GET("/generateCaptcha", admin.AdminController{}.GenerateCaptcha)
		adminrouter.GET("/testSession", admin.AdminController{}.TestSession)
		adminrouter.POST("/upload", admin.AdminController{}.File)

	}
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("I'am middleware!")
	}
}
