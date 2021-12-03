package util

import (
	"github.com/golang-jwt/jwt"
	"go_blog/pkg/setting"
	"time"
)

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(userName string, password string) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 3)
	claims := Claims{userName, password, jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: "go-blo"}}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(setting.Config.App.JwtSecret))
	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.Config.App.JwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
