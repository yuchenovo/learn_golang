package dao

import (
	"context"
	"gorm.io/gorm"
	"study_gin/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (s *UserDao) FindUser(name string) (user *model.User) {
	s.DB.Select("address").Where("username=?", name).First(&user)
	return
}

func (s *UserDao) UpdateUser(name string) (user *model.User) {
	s.DB.Model(&model.User{}).Where("username=?", name).Update("address", "北京市")
	return
}
