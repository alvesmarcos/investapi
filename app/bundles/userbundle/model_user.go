package userbundle

type User struct {
  Id            int     `json:"id"`
  Username      string  `json:"username"`
  Password      string  `json:"password"`
}

func NewUser(username string, password string) *User {
  return &User { Id: 0, Username: username, Password: password }
}

func (user *User) Copy(u *User) {
  user.Username = u.Username
  user.Password = u.Password
}

func (user User) Validate() bool {
  if len(user.Username) == 0 || len(user.Password) == 0 {
    return false
  }
  return true
}
