package main

import (
  "net/http"
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

  // key, err := as.NewKey("test", "set-test", "key-test")
  // if err != nil {
  //   log.Fatal(err)
  // }

  // bins := as.BinMap{
  //   "username": "test",
  //   "password": "test-pwd",
  // }
  //
  // err = client.Put(nil, key, bins)
  // if err != nil {
  //   log.Fatal(err)
  // }
  //
  // rec, err := client.Get(nil, key)
  // if err != nil {
  //   log.Fatal(err)
  // }
  //
  // log.Println(rec)

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

  e.GET("/data/:id", func(c echo.Context) error {
    return controller.GetOneData(c, client)
  })

  e.POST("/data", func(c echo.Context) error {
    return controller.CreateData(c, client)
  })

  // Server
  e.Logger.Fatal(e.Start(":8080"))
}