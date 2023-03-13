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
	//dsn := "root:Strong@01@tcp(172.31.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var user User
	//db.AutoMigrate(&User{}, &CreditCard{})
	//db.First(&user)
	//db.Model(&user).Association("CreditCard").Clear()
	r := gin.Default()
	//引入抽离出来的路由分组
	routes.AdminRouterGroup(r)
	r.Run()
}
