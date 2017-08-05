package userbundle

import "github.com/jinzhu/gorm"

type UserMapperPSQL struct {
  db *gorm.DB
}

func NewUserMapperPSQL(db *gorm.DB) *UserMapperPSQL {
  return &UserMapperPSQL { db: db }
}

func (ump *UserMapperPSQL) FindAll() ([]User, error) {
  var users []User

  ump.db.Find(&users)
  usersx := make([]User, len(users))

  for i, u := range users {
    usersx[i].Id = u.Id
    usersx[i].Username = u.Username
    usersx[i].Password =  "field value omitted"
	}
  return usersx, nil
}

func (ump *UserMapperPSQL) FindUser(u *User) (User, error) {
  var user User

  ump.db.Where("username = ? and password = ?", u.Username, u.Password).First(&user)

  return user, nil
}

func (ump *UserMapperPSQL) Insert(user *User) error {
  return ump.db.Create(user).Error
}

func (ump *UserMapperPSQL) Delete(id int) error {
  return ump.db.Delete(&User {Id: id}).Error
}

func (ump *UserMapperPSQL) Update(user *User) error {
  var u User

  ump.db.First(&u, user.Id)
  u.Copy(user)

  return ump.db.Save(&u).Error
}
