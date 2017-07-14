Learning golang

This is just a learning space to make a Backend API

Response Payload Struct :

```go
type Response struct {
  Errors    error         `json:"error,omitempty"`
  Message   string        `json:"message,omitempty"`
  Data      interface{}   `json:"data,omitempty"`
}
```

Routes :

Routes | Methods | Desc
--- | --- | ---
/users | GET | Get all Users Data
/users/:username | GET | Get 1 User Data
/users/register | POST | Register a user
/users/login | POST | Login a user
/users | PUT | Update 1 user data
/users | DELETE | Delete 1 user data