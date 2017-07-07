package controller

import (
  "net/http"
  "strconv"
  "log"
  "github.com/labstack/echo"
  . "github.com/aerospike/aerospike-client-go"
)

/*

Aerospike      |  SQL         | MongoDB
---------------------------------------------
namespace      |  db          | db
sets           |  table       | collection
key            |  primary_key |
bin            |  column      |
record         |  row         |

*/

type User struct {
  Id        int     `json:"id"`
  Username  string  `json:"username"`
  Password  string  `json:"password"`
  Email     string  `json:"email"`
}

var (
  idSeq = 1
  users = map[int]*User{}

  ErrInternalServer = "Internal Server Error"
)

func panicOnError(err error) {
  if err != nil {
    log.Fatal("Error : %d", err)
    panic(err)
  }
}

// Controller
func GetOneData(c echo.Context, client *Client) error {
  id, _ := strconv.Atoi(c.Param("id"))
  return c.JSON(http.StatusOK, users[id])
}

func CreateData(c echo.Context, client *Client) error {

  key, err := NewKey("test", "Users-test", "user-key-test")
  panicOnError(err)

  u := new(User)
  u.Id = idSeq

  if err := c.Bind(u); err != nil {
    panicOnError(err)
    return c.JSON(http.StatusInternalServerError, ErrInternalServer)
  }

  userBin := BinMap{
    "id"      : u.Id,
    "username": u.Username,
    "password": u.Password,
    "email"   : u.Email,
  }

  err = client.Put(nil, key, userBin)
  panicOnError(err)
  log.Println(u)

  users[u.Id] = u
  idSeq++
  return c.JSON(http.StatusCreated, u)
}

// func UpdateUser(c echo.Context) {
//   d := new(Data)
//
//
// }

