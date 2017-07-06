package main

import (
  "net/http"

  "github.com/andrasat/golang-basic-api/controller"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  // Echo init
  e := echo.New()

  // Middlewares
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // CORS
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
    }))

  // Routes
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello Go 1234567890")
  })
  e.GET("/data/:id", controller.GetOneData)
  e.POST("/data", controller.CreateData)

  // Server
  e.Logger.Fatal(e.Start(":3000"))
}