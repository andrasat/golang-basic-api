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
  Email     string  `json:"email"`
  Password  string  `json:"password"`
}

type Response struct {
  Data      *Recordset `json:"data"`
  Message   error `json:"test-message"`
}

var (
  // users               = map[int]*User{}
  namespace           = "test"
  set                 = "Users-test"
  scPolicy            = NewScanPolicy()

  ErrInternalServer   = "Internal Server Error"
  ErrBadRequest       = "Bad Request"
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
  return c.JSON(http.StatusOK, id)
}

func GetAllUsers(c echo.Context, client *Client) error {
  
  var users []BinMap

  scPolicy := NewScanPolicy()
  scPolicy.ConcurrentNodes = true
  scPolicy.Priority = HIGH
  scPolicy.IncludeBinData = true

  recordSet, err := client.ScanAll(scPolicy, namespace, set)
  // seeRec := <-recordSet.Records // Handle if set is not found inside the namespace -> memory pointer return null

  if err != nil {
    panicOnError(err)
    return c.JSON(http.StatusBadRequest, err)
  }

  for res := range recordSet.Results() {
    if res.Err != nil {
      return c.JSON(http.StatusBadRequest, res.Err)
    } else {
      users = append(users, res.Record.Bins)
    }
  }

  return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context, client *Client) error {

  u := new(User)

  if err := c.Bind(u); err != nil {
    panicOnError(err)
    return c.JSON(http.StatusInternalServerError, ErrInternalServer)
  }

  key, err := NewKey(namespace, set, u.Email)
  panicOnError(err)

  userBin := BinMap{
    "email"   : u.Email,
    "password": u.Password,
  }

  err = client.Put(nil, key, userBin)
  panicOnError(err)
  log.Println(u)


  return c.JSON(http.StatusCreated, u)
}

// func UpdateUser(c echo.Context) {
//   d := new(Data)
//
//
// }

