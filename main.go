package main

import (
  "net/http"
  "log"

  ctrl "github.com/andrasat/golang-basic-api/controller"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  as "github.com/aerospike/aerospike-client-go"
)

func main() {
  // Aerospike
  client, err := as.NewClient("127.0.0.1", 3000)
  if err != nil {
    log.Fatal(err)
  }

  defer client.Close()

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

  e.GET("/users/:email", func(c echo.Context) error {
    return ctrl.GetOneUser(c, client)
  })

  e.GET("/users", func(c echo.Context) error {
    return ctrl.GetAllUsers(c, client)
  })

  e.POST("/users", func(c echo.Context) error {
    return ctrl.CreateUser(c, client)
  })

  // Server
  e.Logger.Fatal(e.Start(":8080"))
}