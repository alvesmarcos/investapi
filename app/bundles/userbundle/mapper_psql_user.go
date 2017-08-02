package userbundle

import (
	"github.com/jinzhu/gorm"
)

type UserMapperPSQL struct {
  db *gorm.DB
}

func NewUserMapperPSQL(db *gorm.DB) *UserMapperPSQL {
  return &UserMapperPSQL { db: db }
}

func (ump *UserMapperPSQL) GetAll() ([]User, error) {
	var users []User

	user := ump.db.Find(&users)
	return users, nil
}

func (ump *UserMapperPSQL) CheckNameAndPasswdMatches(username, password string) (bool, error) {
	var user User

	ump.db.Where("username = ? and password = ?", username, password).First(&user)

	if user.id == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (ump *UserMapperPSQL) Insert(user *User) error {
	return ump.db.Create(user).Error
}

func (ump *UserMapperPSQL) Delete(id uint32) error {
	return ump.db.Delete(&User {Id: id}).Error
}

func (ump *UserMapperPSQL) Update(user *User) error {
	var u User

	ump.db.Where("id = ?", user.id).Find(&u)
	u.Copy(&user)

	return ump.db.Update(&u).Error
}
