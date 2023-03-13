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

	//引入抽离出来的路由分组
	routes.AdminRouterGroup(r)
	r.Run()
}
