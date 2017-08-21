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

  err := ump.db.Find(&users).Error
  usersx := make([]User, len(users))

  if err == nil {
    for i, u := range users {
      usersx[i].Id = u.Id
      usersx[i].Name = u.Name
      usersx[i].Username = u.Username
      usersx[i].Password =  "field value omitted"
    }
  }
  return usersx, err
}

func (ump *UserMapperPSQL) FindUserById(id int) (User, error) {
  var user User

  err := ump.db.First(&user, id).Error

  if err == nil {
    user.Password = "field value omitted"
  }
  return user, err
}

func (ump *UserMapperPSQL) FindUser(u *User) (User, error) {
  var user User

  err := ump.db.Where("username = ? and password = ?", u.Username, u.Password).First(&user).Error

  return user, err
}

func (ump *UserMapperPSQL) Insert(user *User) error {
  return ump.db.Create(user).Error
}

func (ump *UserMapperPSQL) Delete(id int) error {
  return ump.db.Delete(&User {Id: id}).Error
}

func (ump *UserMapperPSQL) Update(user *User) error {
  return ump.db.Save(user).Error
}
