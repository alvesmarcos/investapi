package userbundle

type User struct {
  Id            int     `gorm:"AUTO_INCREMENT" json:"id"`
  Username      string  `gorm:"not null;unique" json:"username"`
  Password      string  `gorm:"not null" json:"password"`
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

func (user *User) CompareAndSwap(u User) {
  if len(user.Username) != len(u.Username) && len(u.Username) > 0 {
    user.Username = u.Username
  }
  if len(user.Password) != len(u.Password) && len(u.Password) > 0 {
    user.Password = u.Password
  }
}
