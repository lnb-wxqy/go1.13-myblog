package common

import (
	"github.com/dgrijalva/jwt-go"
	"myblog/model"
	"time"
)

//密钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//发放token
func ReleaseToken(user *model.User) (string, error) {
	//设置过期时间
	expireTime := time.Now().Add(30 * time.Minute)

	//TODO 测试
	//expireTime := time.Now().Add(5 * time.Second)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(), //token发放时间
			Issuer:    "lnb-wxqy",        //token发放者
			Subject:   "user token",      //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥生成token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
