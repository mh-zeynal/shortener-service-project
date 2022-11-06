package middlewares

import(
	"fmt"
	"net/http"
	
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
	cookie, _ := c.Cookie("token")
	if cookie == nil {
		return c.String(http.StatusUnauthorized, "unauthorized")
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return c.String(http.StatusUnauthorized, "unauthorized")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["usr"], claims["exp"])
	}
	return next(c)
	}
}