package controller

import (
  "net/http"
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
  Errors    error         `json:"error,omitempty"`
  Message   string        `json:"message,omitempty"`
  Data      interface{}   `json:"data,omitempty"`
}

var (
  // users               = map[int]*User{}
  namespace           = "test"
  set                 = "Users-test"

  OKMessage           = "OK"
  ErrInternalServer   = "Internal Server Error"
  ErrBadRequest       = "Bad Request"
  RecNotFound         = "Record Not Found"
)

func panicOnError(err error) {
  if err != nil {
    log.Fatal("Error : %d", err)
    panic(err)
  }
}

// Controller
func GetOneUser(c echo.Context, client *Client) error {

  r := new(Response)

  key, err := NewKey(namespace, set, c.Param("email"))
  panicOnError(err)

  rec, err := client.Get(nil, key)
  panicOnError(err)

  if rec == nil {
    r.Message = RecNotFound
    return c.JSON(http.StatusNotFound, r)
  }

  r.Data = rec.Bins
  return c.JSON(http.StatusOK, r)
}

func GetAllUsers(c echo.Context, client *Client) error {

  r := new(Response)
  var users []BinMap

  recordSet, err := client.ScanAll(nil, namespace, set)
  // seeRec := <-recordSet.Records // Handle if set is not found inside the namespace -> memory pointer return null

  if err != nil {
    r.Errors = err
    panicOnError(err)
    return c.JSON(http.StatusBadRequest, r)
  }
  // if <-recordSet.Records == nil {
  //   r.Message = RecNotFound
  //   return c.JSON(http.StatusNotFound, r)
  // }
  // err = recordSet.Close();
  // panicOnError(err)

  for res := range recordSet.Results() {
    if res.Err != nil {
      r.Errors = err
      return c.JSON(http.StatusBadRequest, r)
    } else {
      users = append(users, res.Record.Bins)
    }
  }

  r.Data = users
  return c.JSON(http.StatusOK, r)
}

func CreateUser(c echo.Context, client *Client) error {

  u := new(User)
  r := new(Response)

  if err := c.Bind(u); err != nil {
    panicOnError(err)
    r.Errors = err
    r.Message = ErrInternalServer
    return c.JSON(http.StatusInternalServerError, r)
  }

  key, err := NewKey(namespace, set, u.Email)
  panicOnError(err)

  userBin := BinMap{
    "email"   : u.Email,
    "password": u.Password,
  }

  err = client.Put(nil, key, userBin)
  panicOnError(err)

  r.Message = OKMessage
  r.Data = u
  return c.JSON(http.StatusCreated, r)
}

func UpdateUser(c echo.Context, client *Client) error {

  r := new(Response)
  u := new(User)

  key, err := NewKey(namespace, set, c.Param("email"))
  panicOnError(err)

  rec, err := client.Get(nil, key)
  panicOnError(err)

  if rec == nil {
    r.Message = RecNotFound
    return c.JSON(http.StatusNotFound, r)
  }

  u.Email = rec.Bins["email"].(string)
  u.Password = rec.Bins["password"].(string)

  if err := c.Bind(u); err != nil {
    panicOnError(err)
    r.Errors = err
    r.Message = ErrInternalServer
    return c.JSON(http.StatusInternalServerError, r)
  }

  userBin := BinMap{
    "email"     : u.Email,
    "password"  : u.Password,
  }

  err = client.Put(nil, key, userBin)
  panicOnError(err)

  r.Data = u
  r.Message = OKMessage
  return c.JSON(http.StatusOK, r)
}

