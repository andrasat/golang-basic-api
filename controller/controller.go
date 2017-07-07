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
func GetOneUser(c echo.Context, client *Client) error {
  id, _ := strconv.Atoi(c.Param("id"))
  return c.JSON(http.StatusOK, users[id])
}

func GetAllUsers(c echo.Context, client *Client) error {

  recordSet, err := client.ScanAll(nil, "test", "Users-test")
  panicOnError(err)

  for res := range recordSet.Results() {
    if res.Err != nil {
      panicOnError(res.Err)
    } else {
      log.Println(res.Record)
      return c.JSON(http.StatusOK, res.Record)
    }
  }

  return c.JSON(http.StatusInternalServerError, ErrInternalServer)
}

func CreateUser(c echo.Context, client *Client) error {

  u := new(User)
  u.Id = idSeq

  if err := c.Bind(u); err != nil {
    panicOnError(err)
    return c.JSON(http.StatusInternalServerError, ErrInternalServer)
  }

  key, err := NewKey("test", "Users-test", u.Username)
  panicOnError(err)

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

