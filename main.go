package main

import (
  "log"

  "github.com/andrasat/golang-basic-api/controller"
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
  //e.Use(middleware.Recover())

  // CORS
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
  }))

  ctr := &controller.Controller{DB: client}

  // Routes
  e.GET("/users/:username", ctr.GetOneUser)
  e.GET("/users", ctr.GetAllUsers)
  e.POST("/users/register", ctr.CreateUser)
  e.POST("/users/login", ctr.LoginUser)
  e.PUT("/users/:username", ctr.UpdateUser)

  // Server
  e.Logger.Fatal(e.Start(":8080"))
}