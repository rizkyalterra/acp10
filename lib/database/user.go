package database

import (
	"acp10/configs"
	"acp10/models/news"
)

func GetDataUserAll() (dataResult []news.News, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}
