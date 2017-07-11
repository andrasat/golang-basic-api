package controller

import (
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

type Controller struct {
  DB *Client
}

type Response struct {
  Errors    error         `json:"error,omitempty"`
  Message   string        `json:"message,omitempty"`
  Data      interface{}   `json:"data,omitempty"`
}

const (
  namespace           = "test"
  set                 = "Users-test"

  OKMessage           = "OK"
  FAILMessage         = "Fail"
  ErrInternalServer   = "Internal Server Error"
  ErrBadRequest       = "Bad Request"
  RecNotFound         = "Record Not Found"
  NotAuthorized       = "Wrong Password"
)
