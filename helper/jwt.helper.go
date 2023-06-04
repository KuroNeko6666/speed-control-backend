package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/KuroNeko6666/speed-control-backend.git/database/model"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user model.User, key string) (string, error) {
	keyToken := []byte(key)
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(10 * time.Hour).Unix(),
	})
	token, err := rawToken.SignedString(keyToken)

	return token, err
}

func GetJWTFromHeader(data map[string]string) string {
	return strings.Split(data["Authorization"], " ")[1]
}

func ExtractClaims(tokenStr string, key string) (jwt.MapClaims, error) {
	secret := []byte(key)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("error extract token")
	}
}

func ExtractUserFromJWT(tkn string, key string) (model.User, error) {
	var user model.User

	claims, err := ExtractClaims(tkn, key)

	if err != nil {
		return user, err
	}

	user.ID = claims["id"].(string)
	user.Username = claims["username"].(string)
	user.Email = claims["email"].(string)

	return user, nil
}

func ExtractExpirerFromJWT(tkn string, key string) (time.Time, error) {
	var date time.Time
	layout := "2006-01-02 15:04:05"

	claims, err := ExtractClaims(tkn, key)

	if err != nil {
		return date, err
	}

	date, err = time.Parse(layout, claims["exp"].(string))

	if err != nil {
		return date, err
	}

	return date, nil

}

func ExtractUserFromHeader(data map[string]string, key string) (model.User, error) {
	var user model.User
	tkn := GetJWTFromHeader(data)
	claims, err := ExtractClaims(tkn, key)

	if err != nil {
		return user, err
	}

	user.ID = claims["id"].(string)
	user.Username = claims["username"].(string)
	user.Email = claims["email"].(string)

	return user, nil
}
