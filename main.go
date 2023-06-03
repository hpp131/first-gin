package main

import (
	"first-gin/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"size:100"`
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID int
}

func main() {
	r := gin.Default()
	//r.MaxMultipartMemory = 8 << 20
	//Invoke sessions dependency and use cookie as stroage engine
	//store := cookie.NewStore([]byte("secret123"))
	// import sessions as middleware
	//r.Use(sessions.Sessions("mySession", store))

	r1 := r.Group("/admin")
	{
		r1.GET("/test", func(context *gin.Context) {
			context.String(200, "success")
		})
	}
	//引入抽离出来的路由分组
	routes.AdminRouterGroup(r)
	r.Run()
}
