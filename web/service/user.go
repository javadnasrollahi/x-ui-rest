package service

import (
	"errors"
	"x-ui/database"
	"x-ui/database/model"
	"x-ui/logger"

	"gorm.io/gorm"
)

type UserService struct {
}

func (s *UserService) GetFirstUser() (*model.User, error) {
	db := database.GetDB()

	user := &model.User{}
	err := db.Model(model.User{}).
		First(user).
		Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CheckUser(username string, password string) *model.User {
	db := database.GetDB()

	user := &model.User{}
	err := db.Model(model.User{}).
		Where("username = ? and password = ?", username, password).
		First(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		logger.Warning("check user err:", err)
		return nil
	}
	return user
}

func (s *UserService) UpdateUser(id int, username string, password string, apitoken string) error {
	db := database.GetDB()
	return db.Model(model.User{}).
		Where("id = ?", id).
		Update("username", username).
		Update("password", password).
		Update("apitoken", apitoken).
		Error
}

func (s *UserService) UpdateFirstUser(username string, password string, apitoken string) error {
	if username == "" {
		return errors.New("username can not be empty")
	} else if password == "" {
		return errors.New("password can not be empty")
	} else if apitoken == "" {
		return errors.New("apitoken can not be empty")
	}
	db := database.GetDB()
	user := &model.User{}
	err := db.Model(model.User{}).First(user).Error
	if database.IsNotFound(err) {
		user.Username = username
		user.Password = password
		user.ApiToken = apitoken
		return db.Model(model.User{}).Create(user).Error
	} else if err != nil {
		return err
	}
	user.Username = username
	user.Password = password
	user.ApiToken = apitoken
	return db.Save(user).Error
}
