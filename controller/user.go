package controller

import (
  "net/http"

  bc "golang.org/x/crypto/bcrypt"
  "github.com/labstack/echo"
  . "github.com/aerospike/aerospike-client-go"
)

type User struct {
  Username  string  `json:"username"`
  Email     string  `json:"email,omitempty"`
  Token     string  `json:"token,omitempty"`
  Password  string  `json:"password,omitempty"`
}

/*  GET ONE USER
    --------------------------------------------------
*/

func (ct *Controller) GetOneUser(c echo.Context) error {

  r := new(Response)

  key, err := NewKey(namespace, set, c.Param("username"))
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  rec, err := ct.DB.Get(nil, key)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  if rec == nil {
    r.Message = RecNotFound
    return c.JSON(http.StatusNotFound, r)
  }

  r.Data = rec.Bins
  return c.JSON(http.StatusOK, r)
}

/*  GET ALL USERS
    --------------------------------------------------
*/

func (ct *Controller) GetAllUsers(c echo.Context) error {

  r := new(Response)
  var users []BinMap

  recordSet, err := ct.DB.ScanAll(nil, namespace, set)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  // If set is not found inside the namespace -> memory pointer return null
  // if <-recordSet.Records == nil {
  //   r.Message = RecNotFound
  //   return c.JSON(http.StatusNotFound, r)
  // }
  // err = recordSet.Close();

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

/*  REGISTER ONE USER
    --------------------------------------------------
*/

func (ct *Controller) CreateUser(c echo.Context) error {

  u := new(User)
  r := new(Response)

  if err := c.Bind(u); err != nil {
    r.Errors = err
    r.Message = ErrInternalServer
    return c.JSON(http.StatusInternalServerError, r)
  }

  key, err := NewKey(namespace, set, u.Username)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  hashedPass, err := bc.GenerateFromPassword([]byte(u.Password), bc.DefaultCost)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }
  u.Password = string(hashedPass)

  userBin := BinMap{
    "username": u.Username,
    "email"   : u.Email,
    "password": hashedPass,
  }

  err = ct.DB.Put(nil, key, userBin)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  r.Message = OKMessage
  r.Data = u
  return c.JSON(http.StatusCreated, r)
}

/*  LOGIN ONE USER
    --------------------------------------------------
*/

func (ct *Controller) LoginUser(c echo.Context) error {

  r := new(Response)
  u := new(User)

  if err := c.Bind(u); err != nil {
    r.Errors = err
    r.Message = ErrInternalServer
    return c.JSON(http.StatusInternalServerError, r)
  }

  key, err := NewKey(namespace, set, u.Username)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  rec, err := ct.DB.Get(nil, key)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  } else if rec == nil {
    r.Message = RecNotFound
    return c.JSON(http.StatusNotFound, r)
  }

  return c.JSON(http.StatusOK, "Test")
}

/*  UPDATE ONE USER
    --------------------------------------------------
*/

func (ct *Controller) UpdateUser(c echo.Context) error {

  r := new(Response)
  u := new(User)

  key, err := NewKey(namespace, set, c.Param("username"))
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  rec, err := ct.DB.Get(nil, key)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  if rec == nil {
    r.Message = RecNotFound
    return c.JSON(http.StatusNotFound, r)
  }

  u.Email = rec.Bins["email"].(string)
  u.Password = rec.Bins["password"].(string)

  if err := c.Bind(u); err != nil {
    r.Errors = err
    r.Message = ErrInternalServer
    return c.JSON(http.StatusInternalServerError, r)
  }

  userBin := BinMap{
    "email"     : u.Email,
    "password"  : u.Password,
  }

  err = ct.DB.Put(nil, key, userBin)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  r.Data = u
  r.Message = OKMessage
  return c.JSON(http.StatusOK, r)
}