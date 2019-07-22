package model

import "pipiao.yguilai.com/database/mysql"

type Auth struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

func CheckAuth(appid, secret string) bool {
	var auth Auth
	db := mysql.GetDb()
	db.Select("id").Where(Auth{Appid: appid, Secret: secret}).First(&auth)
	if auth.Id > 0 {
		return true
	}
	return false
}
