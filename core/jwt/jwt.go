package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"nimbus/core/config"
	"fmt"
)

type myClaims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

const EXPIRES_AT = 10080

func CreateJWT(data interface{}) (token string, error error) {
	appConfig := config.GetAppConfig().Config
	secretJWT := []byte(appConfig["SECRET_JWT"])

	claims := myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: EXPIRES_AT,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secretJWT)
	return token, err
}

func ParseJWT(tokenString string) (data interface{}, err error) {

	appConfig := config.GetAppConfig().Config
	secretJWT := []byte(appConfig["SECRET_JWT"])

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretJWT, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
