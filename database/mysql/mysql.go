package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

const (
	_DRIVER   = "mysql"
	_USERNAME = "root"
	_PASSWORD = "password"
	_HOST     = "localhost"
	_PORT     = "3306"
	_DBNAME   = "dbname"
	_CHARSET  = "utf8"
)

// 设置数据库连接池
func init() {
	var err error
	url := _USERNAME + ":" + _PASSWORD + "@tcp(" + _HOST + ":" + _PORT + ")/" + _DBNAME + "?charset=" + _CHARSET
	db, err = gorm.Open(_DRIVER, url)
	if err != nil {
		log.Panic(err)
		return
	}
	// 连接池最大闲置连接数
	db.DB().SetMaxIdleConns(100)
	// 连接池最大打开连接数
	db.DB().SetMaxOpenConns(150)
	// 连接可被重新使用的最大时间间隔
	db.DB().SetConnMaxLifetime(time.Second * 30)
	// gin表名默认为复数修改为单数 即不加s
	db.SingularTable(true)

	// 给默认表名添加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "ppa_" + defaultTableName
	}
}

func GetDb() *gorm.DB {
	return db
}
