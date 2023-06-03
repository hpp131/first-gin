package admin

import (
	"first-gin/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type AdminController struct {
	//通过结构体的嵌套实现controller的继承
	BaseController
}

// 调用models中的公共方法
var timeNow = models.UnixToDate(time.Now().Unix())

func (c AdminController) Index(ctx *gin.Context) {
	//ctx.String(200, "admin-index api success")
	c.Success(ctx)
}

func (c AdminController) User(ctx *gin.Context) {
	name, _ := ctx.Get("name")
	result, _ := name.(string)
	ctx.String(200, result+timeNow)
}

func (c AdminController) GenerateCaptcha(ctx *gin.Context) {
	id, b64s, err := models.Generate()
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaid":   id,
		"captchadata": b64s,
	})
}

func (c AdminController) TestSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("username", "zhangsan")
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"sessionid": session.Get("username"),
	})
}

func (c AdminController) Date(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"date": time.Now(),
	})
}

func (c AdminController) File(ctx *gin.Context) {
	file, _ := ctx.FormFile("myfile")
	log.Println(file.Filename)
	ctx.SaveUploadedFile(file, file.Filename)
	ctx.String(http.StatusOK, fmt.Sprintf("%s upload successfully", file.Filename))
}
