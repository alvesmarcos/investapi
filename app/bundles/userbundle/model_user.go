package userbundle

type User struct {
  Id            uint32  `json:"id"`
  Username      string  `json:"username"`
  Password      string  `json:"password"`
}

func NewUser(id uint32, username, password string) *User {
  return &User { Id: id, Username: username, Password: password }
}

func (user *User) Copy(u *User) {
  user.Username = u.Username
  user.Password = u.Password
}
