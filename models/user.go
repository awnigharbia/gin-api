package models

import (
	"errors"

	"github.com/go-xorm/xorm"
)

type User struct {
	ID int `json:"id" xorm:"id"`
	Name string `json:"name"`
	Email string `json:"email" xorm:"email"`
	Password string `json:"password" xorm:"password"`
	FcmToken string `json:"fcm_token" xorm:"fcm_token"`
}

func GetUserByID(db *xorm.Engine, id string) (User, error) {
	var user User
	has, err := db.Table("users").Where("id = ?", id).Get(&user)

	if err != nil {
		return User{}, err
	}
	if !has {
		return User{}, errors.New("user not found")
	}

	return user, nil
}
