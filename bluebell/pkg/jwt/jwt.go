package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
@author RandySun
@create 2022-01-15-13:23
*/

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("冬天冬天悄悄过去")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(UserId int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		UserId,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")

}
