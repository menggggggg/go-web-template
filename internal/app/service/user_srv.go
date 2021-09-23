package service

import (
	"context"

	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/dao"
	"github.com/menggggggg/go-web-template/internal/app/schema"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSrv), "*"))

type UserSrv struct {
	UserDao *dao.UserDao
}

func (a *UserSrv) Get(ctx context.Context, name string) (*schema.User, error) {
	return a.UserDao.Get(ctx, name)
}
