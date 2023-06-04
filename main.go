package main

import (
	"first-gin/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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
	//使用第三方库做为middleware实现session功能
	//store := cookie.NewStore([]byte("mysecret"))
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("mysecret"))
	r.Use(sessions.Sessions("mySession", store))

	r1 := r.Group("/admin")
	{
		r1.GET("/test", func(context *gin.Context) {
			context.String(200, "success")
		})
		////单独使用cookie
		//r1.GET("/setCookie", func(context *gin.Context) {
		//	context.SetCookie("mycookie_name", "测试", 3600, "/", "localhost", false, true)
		//	context.JSON(200, gin.H{
		//		"settingCookie": true,
		//	})
		//})
		//r1.GET("/getCookie", func(context *gin.Context) {
		//	cookieValue, err := context.Cookie("mycookie_name")
		//	if err != nil {
		//		context.JSON(200, gin.H{
		//			"message": fmt.Sprintf("%s is not found", cookieValue),
		//		})
		//		return
		//	}
		//	context.JSON(200, gin.H{
		//		"cookieValue": cookieValue,
		//	})
		//
		//})
		r1.GET("/hello", func(context *gin.Context) {
			session := sessions.Default(context)
			var count int
			v := session.Get("count")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count++
			}
			session.Set("count", count)
			session.Save()
			context.JSON(200, gin.H{
				"count": session.Get("count"),
			})

		})
	}
	//引入抽离出来的路由分组
	routes.AdminRouterGroup(r)
	r.Run()
}
