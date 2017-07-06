package controller

import (
  "net/http"
  "strconv"
  "log"
  "github.com/labstack/echo"
  "github.com/aerospike/aerospike-client-go"
)

/*

Aerospike      |  SQL
--------------------------------
namespace      |  db
sets           |  table
bin            |  column
key            |  primary_key
record         |  row

*/

type Data struct {
  ID    int     `json:"id"`
  Body  string  `json:"body"`
}

client, err := as.NewClient("127.0.0.1", 3000)
if err != nil {
  log.Fatal(err)
}

key, err := as.NewKey("namespace-test", "set-test", "key-test")
if err != nil {
  log.Fatal(err)
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
  bin := as.NewBin("")

  if err := c.Bind(d); err != nil {
    return c.JSON(http.StatusInternalServerError, ErrInternalServer)
  }

  d.ID = startSequence
  d.Body = c.FormValue("body")

  startSequence++
  return c.JSON(http.StatusCreated, d)
}

func UpdateUser(c echo.Context) {
  d := new(Data)


}

