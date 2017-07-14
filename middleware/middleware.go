package middleware

import (
  "fmt"
  "strings"
  "net/http"

  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"
)

type Middleware struct {

}

func (md *Middleware) JWTAuthenticator(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // fmt.Printf("Authorization Header :  %v\n\n", c.Request().Header["Authorization"][0])
    // fmt.Printf("Token String :  %v\n\n", tokenString[1])

    tokenString := strings.Split(c.Request().Header["Authorization"][0], " ")

    token, err := jwt.Parse(tokenString[1], func (token *jwt.Token) (interface{}, error) {

      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
      }

      return []byte("SECRET"), nil
    })
    if err != nil {
      fmt.Printf("The Error : %v\n\n", err)
      return c.JSON(http.StatusInternalServerError, err)
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      c.Set("username", claims["username"])
      c.Set("exp", claims["exp"])
      return next(c)
    }

    return c.JSON(http.StatusUnauthorized, "Not Authorized")
  }
}