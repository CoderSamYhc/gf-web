package service

import (
	"context"
	"gf-web/internal/model/entity"
	"gf-web/internal/service/internal/dao"
)

type sUser struct {

}

func UserService() *sUser {
	return &sUser{}
}

func (s *sUser) Info(ctx context.Context, id int) (*entity.User, error) {
	var (
		userInfo *entity.User
		err error
	)
	if err = dao.User.Ctx(ctx).WherePri(id).Scan(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

