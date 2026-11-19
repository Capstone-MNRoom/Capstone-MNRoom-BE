package middlewares

import (
	_constants "be9/mnroom/constants"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(_constants.SECRET_JWT),
	})
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userId,
		"exp":        time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(_constants.SECRET_JWT))
}

func ExtractToken(e echo.Context) (int, error) {
	user, ok := e.Get("user").(*jwt.Token)
	if !ok || user == nil {
		return 0, fmt.Errorf("token not found")
	}
	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		if userId, ok := claims["userId"].(float64); ok {
			return int(userId), nil
		}
	}
	return 0, fmt.Errorf("token invalid")
}
