package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	Name string
	Age  int
	Job  string
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
	//
	dsn := "root:Strong@01@tcp(172.31.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{})

	r.Run()
}
