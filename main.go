package main

import (
  "net/http"
  "strconv"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

type Data struct {
  ID    int     `json:"id"`
  body  string  `json:"body"`
}

type Error struct {
  Status      int    `json:"status"`
  Title       string `json:"title"`
  Description string `json"description,omitempty`
}

type Response struct {
  Errors []Error `json:"errors,omitempty"`
  Body interface{} `json:"data,omitempty"`
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
  d := new(Data)
  d.Type = "data"

  if err := c.Bind(u); err != nil {
    return c.JSON(http.StatusInternalServerError, makeErrorResponse(err, http.StatusInternalServerError))
  }


  startSequence++
  return c.JSON(http.StatusCreated, d)
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
  e.GET("/data/:id", getOneData)
  e.POST("/data", createData)

  // Server
  e.Logger.Fatal(e.Start(":3000"))
}