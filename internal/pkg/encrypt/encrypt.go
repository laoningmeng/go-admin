package encrypt

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Md5(str string) string {
	md5Sum := md5.Sum([]byte(str))
	md5Str := fmt.Sprintf("%x", md5Sum)
	return md5Str
}

var jwtSecret []byte

type Claims struct {
	Field map[string]string `json:"field"`
	jwt.StandardClaims
}

func GenerateToken(field map[string]string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		field,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-admin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
