package api

import (
	"api_gin/model"
	e "api_gin/pkg/e"
	"api_gin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type auth struct {
	Appid  string `valid:"Required; MaxSize(255)"`
	Secret string `valid:"Required; MaxSize(255)"`
}

// 获取token Api
// Get /auth?appid=xxx&secret=xxx
func GetAuth(c *gin.Context) {
	appid := c.Query("appid")
	secret := c.Query("secret")
	valid := validation.Validation{}
	a := auth{Appid: appid, Secret: secret}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := model.CheckAuth(appid, secret)
		if isExist {
			// 生存token
			token, err := util.GenerateToken(appid, secret)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
				log.Println(err)
			} else {
				// 添加到header里的参数名称
				data["tokenType"] = "Authentication"
				// token有效期
				data["expireDur"] = time.Now().Add(time.Hour*24*7*2).Unix() - time.Now().Unix()
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
