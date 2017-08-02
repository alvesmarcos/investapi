package userbundle

// User struct
type User struct {
  Id        uint32  `json:"id"`
  Name      string  `json:"name"`
  Password  string  `json:"name"`
}

// NewUser create a new User
func NewUser(id uint32, name, password string) *User {
  return &User { Id: id, Name: name, Password: password }
}
