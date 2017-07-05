package main

import (
  "net/http"
  "strconv"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

type Data struct {
  ID int  `json:"id"`
  body string `json:"body"`
}

var (
  startSequence = 1
  datas = map[int]*Data{}
)

// Controller
func getOneData(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  return c.JSON(http.StatusOK, datas[id])
}

func createData(c echo.Context) error {
  u := &Data{
    ID: startSequence,
  }
  if err := c.Bind(u); err != nil {
    return err
  }
  datas[u.ID] = u
  startSequence++
  return c.JSON(http.StatusCreated, u)
}

/*
func updateUser(c echo.Context) {

}
*/

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

  // Server
  e.Logger.Fatal(e.Start(":3000"))
}