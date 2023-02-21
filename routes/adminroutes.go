package routes

import (
	"first-gin/controller/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouterGroup(r *gin.Engine) {
	adminrouter := r.Group("/admin")
	{
		adminrouter.GET("/index", func(context *gin.Context) {
			context.String(200, "admin-index api success")
		})
		adminrouter.GET("/user", admin.AdminController{}.User)
	}
}
