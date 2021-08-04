package database

import (
	"acp10/configs"
	"acp10/models/news"
	"acp10/models/users"
)

func GetDataUserAll() (dataResult []news.News, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func LoginUser(userLogin users.UserLogin) (users.User, error) {
	var userDB users.User
	err := configs.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).First(&userDB).Error
	if err != nil {
		return userDB, err
	}
	return userDB, nil
}
