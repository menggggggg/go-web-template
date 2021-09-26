package service

import (
	"context"

	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/dao"
	"github.com/menggggggg/go-web-template/internal/app/schema/user"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSrv), "*"))

type UserSrv struct {
	UserDao *dao.UserDao
}

func (a *UserSrv) Create(ctx context.Context, request *user.CreateRequest) (*user.User, error) {
	return nil, nil
}

func (a *UserSrv) GetByName(ctx context.Context, request *user.GetByNameRequest) (*user.User, error) {
	return nil, nil
}

func (a *UserSrv) Get(ctx context.Context, request *user.GetRequest) (*user.User, error) {
	return nil, nil
}

func (a *UserSrv) Update(ctx context.Context, request *user.UpdateRequest) (*user.User, error) {
	return nil, nil
}

func (a *UserSrv) Delete(ctx context.Context, request *user.DeleteRequest) error {
	return nil
}
