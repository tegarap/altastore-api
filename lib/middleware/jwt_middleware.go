package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

func CreateToken(id int, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = id
	claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractToken(c echo.Context) (int, bool) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		isAdmin := claims["isAdmin"]
		return int(userId), isAdmin.(bool)
	}
	return 0, false
}

//==============================================================
//	MIDDLEWARE FOR UNIT TESTING
//==============================================================

func CreateTokenTest(id int, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = id
	claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT_TEST")))
}

func ExtractTokenTest(c echo.Context) (int, bool) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		isAdmin := claims["isAdmin"]
		return int(userId), isAdmin.(bool)
	}
	return 0, false
}