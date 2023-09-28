package service

import (
	"context"
	"study_gin/dao"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) GetNameService(ctx context.Context, param string) string {
	userDao := dao.NewUserDao(ctx)
	name := userDao.FindUser(param)
	return name.Address
}
func (s *UserSrv) UpdateUserService(ctx context.Context, name string) (req interface{}) {
	userDao := dao.NewUserDao(ctx)
	userDao.UpdateUser(name)
	return
}
