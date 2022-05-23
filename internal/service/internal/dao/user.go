package dao

import "gf-web/internal/service/internal"

type userDao struct {
	*internal.UserDao
}

var User = userDao{
	internal.NewUserDao(),
}
