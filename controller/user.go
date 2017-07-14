package controller

import (
  "net/http"
  "time"
  "fmt"

  bc "golang.org/x/crypto/bcrypt"
  "github.com/dgrijalva/jwt-go"
  "github.com/labstack/echo"
  as "github.com/aerospike/aerospike-client-go"
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

  key, err := as.NewKey(namespace, set, c.Param("username"))
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

  r.Message = OKMessage
  r.Data = rec.Bins
  return c.JSON(http.StatusOK, r)
}

/*  GET ALL USERS
    --------------------------------------------------
*/

func (ct *Controller) GetAllUsers(c echo.Context) error {

  r := new(Response)
  var users []as.BinMap

  recordSet, err := ct.DB.ScanAll(nil, namespace, set)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  for res := range recordSet.Results() {
    if res.Err != nil {
      r.Errors = err
      return c.JSON(http.StatusBadRequest, r)
    } else {
      users = append(users, res.Record.Bins)
    }
  }

  r.Message = OKMessage
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

  if u.Username == "" || u.Email == "" || u.Password == "" {
    r.Message = FieldCannotEmpty
    return c.JSON(http.StatusBadRequest, r)
  }

  key, err := as.NewKey(namespace, set, u.Username)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  hashedPass, err := bc.GenerateFromPassword([]byte(u.Password), bc.DefaultCost)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  userBin := as.BinMap{
    "username": u.Username,
    "email"   : u.Email,
    "password": string(hashedPass),
  }

  err = ct.DB.Put(nil, key, userBin)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  r.Message = OKMessage
  u.Password = ""
  u.Email = ""
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
    r.Message = ErrInternalServer+", Bind Error"
    return c.JSON(http.StatusInternalServerError, r)
  }

  key, err := as.NewKey(namespace, set, u.Username)
  if err != nil {
    r.Errors = err
    r.Message = ErrInternalServer+", NewKey Error"
    return c.JSON(http.StatusInternalServerError, r)
  }

  rec, err := ct.DB.Get(nil, key)
  if err != nil {
    r.Message = ErrInternalServer+", Aerospike Error"
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  } else if rec == nil {
    r.Message = RecNotFound
    return c.JSON(http.StatusNotFound, r)
  }

  getPass := rec.Bins["password"].(string)

  if err := bc.CompareHashAndPassword([]byte(getPass), []byte(u.Password)); err == nil {

    // Hide Password
    u.Password = ""

    // Create Token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "username"  : u.Username,
      "exp"       : time.Now().Add(time.Hour * 24).Unix(),
    })
    // "iat"       : time.Now().Unix(),
    // "nbf"       : time.Now().Add(time.Hour * 48).Unix(),

    t, err := token.SignedString([]byte("SECRET"))
    if err != nil {
      r.Message = FAILMessage+", Token not generated"
      r.Errors = err
      return c.JSON(http.StatusInternalServerError, r)
    }

    u.Token = t
    r.Message = OKMessage
    r.Data = u
    fmt.Printf("This is your TOKEN : %v\n\n", t)
    return c.JSON(http.StatusOK, r)
  }

  r.Errors = err
  r.Message = NotAuthorized
  return c.JSON(http.StatusUnauthorized, NotAuthorized)
}

/*  UPDATE ONE USER
    --------------------------------------------------
*/

func (ct *Controller) UpdateUser(c echo.Context) error {

  r := new(Response)
  u := new(User)

  key, err := as.NewKey(namespace, set, c.Param("username"))
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

  if err := c.Bind(u); err != nil {
    r.Errors = err
    r.Message = ErrInternalServer
    return c.JSON(http.StatusInternalServerError, r)
  }

  if c.FormValue("Email") == "" {
    r.Message = FieldCannotEmpty
    return c.JSON(http.StatusBadRequest, r)
  }

  if u.Password == "" {
    mailBin := as.NewBin("email", u.Email)

    err = ct.DB.PutBins(nil, key, mailBin)
    if err != nil {
      r.Errors = err
      return c.JSON(http.StatusInternalServerError, r)
    }

    r.Data = u
    r.Message = OKMessage
    return c.JSON(http.StatusOK, r)
  }

  hashedPass, err := bc.GenerateFromPassword([]byte(u.Password), bc.DefaultCost)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  u.Password = string(hashedPass)
  mailBin := as.NewBin("email", u.Email)
  passBin := as.NewBin("password", string(hashedPass))

  err = ct.DB.PutBins(nil, key, mailBin, passBin)
  if err != nil {
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  }

  r.Data = u
  r.Message = OKMessage
  return c.JSON(http.StatusOK, r)
}

/*  DELETE ONE USER
    --------------------------------------------------
*/

func (ct *Controller) DeleteUser(c echo.Context) error {

  r := new(Response)

  uname := c.Get("username").(string)
  // exp := c.Get("exp").(int64)

  key, _ := as.NewKey(namespace, set, uname)

  exists, err := ct.DB.Delete(nil, key)
  if err != nil {
    r.Message = FAILMessage
    r.Errors = err
    return c.JSON(http.StatusInternalServerError, r)
  } else if !exists {
    r.Message = RecNotFound
    r.Data    = exists
    return c.JSON(http.StatusNotFound, r)
  }

  r.Message = OKMessage
  r.Data    = exists
  return c.JSON(http.StatusOK, r)
}