package routes

import (
	"first-gin/controller/admin"
	"first-gin/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouterGroup(r *gin.Engine) {
	// 在路由组中配置中间件
	adminrouter := r.Group("/admin", middleware.AdminMiddleWare)
	{
		adminrouter.GET("/index", func(context *gin.Context) {
			context.String(200, "admin-index api success")
		})
		adminrouter.GET("/user", admin.AdminController{}.User)
	}
}
