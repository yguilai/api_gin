package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 此处为jwt服务端私钥
var jwtSecret []byte = []byte("123456")

type Claims struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
	jwt.StandardClaims
}

// 生成jwt token
// @param appid string
// @param secret string
// @return token string
// @return err error
func GenerateToken(appid, secret string) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24 * 7 * 2)

	claims := Claims{
		appid,
		secret,
		jwt.StandardClaims{
			// token过期时间 此处为2周
			ExpiresAt: expireTime.Unix(),
			// Issuer: 即token发行人, 可随意修改
			Issuer: "ygl",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析token
// @param token string
// @return claims *Claims
// @return err error
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
