package ext

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

// 带权限创建令牌
func GenerateToken(user interface{}) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user,
		"exp":     time.Now().Add(time.Hour * 480).Unix(),
	})

	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("mykey"))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// ParseToken parse JWT token in http header.
func ParseToken(authString string) (t jwt.MapClaims, err error) {

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return nil, errors.New("token error")
	}
	tokenString := kv[1]
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token error")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token error")
	}

}

//func ParseToken(token string) (*Claims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//
//	return nil, err
//}
