package controller

import (
  "net/http"
  "strconv"
  "log"
  "github.com/labstack/echo"
)

type Data struct {
  ID    int     `json:"id"`
  Body  string  `json:"body"`
}
/*
type Error struct {
  Status      int    `json:"status"`
  Title       string `json:"title"`
  Description string `json"description,omitempty`
}

type Response struct {
  Errors []Error `json:"errors,omitempty"`
  Body interface{} `json:"data,omitempty"`
}
*/
var (
  startSequence = 1
  datas = map[int]*Data{}

  ErrInternalServer = "Internal Server Error"
)

// Controller
func GetOneData(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  return c.JSON(http.StatusOK, datas[id])
}

func CreateData(c echo.Context) error {
  d := new(Data)

  if err := c.Bind(d); err != nil {
    return c.JSON(http.StatusInternalServerError, ErrInternalServer)
  }

  d.ID = startSequence
  d.Body = c.FormValue("body")

  startSequence++
  return c.JSON(http.StatusCreated, d)
}

/*
func UpdateUser(c echo.Context) {

}
*/
