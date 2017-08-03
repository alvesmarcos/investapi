package userbundle

import "github.com/jinzhu/gorm"

type UserMapperPSQL struct {
  db *gorm.DB
}

func NewUserMapperPSQL(db *gorm.DB) *UserMapperPSQL {
  return &UserMapperPSQL { db: db }
}

func (ump *UserMapperPSQL) FindAll() ([]string, error) {
	var users []User

	ump.db.Find(&users)
	usernames := make([]string, len(users))

	for i, u := range users {
		usernames[i] = u.Username
	}
	return usernames, nil
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

	ump.db.Where("id = ?", user.Id).Find(&u)
	u.Copy(user)

	return ump.db.Update(&u).Error
}
