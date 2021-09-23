package dao

import (
	"context"

	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/schema"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

// UserRepo 用户存储
type UserDao struct {
	DB *gorm.DB
}

func (a *UserDao) Get(ctx context.Context, name string) (*schema.User, error) {
	// var item User
	// ok, err := util.FindOne(ctx, GetUserDB(ctx, a.DB).Where("id=?", id), &item)
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// } else if !ok {
	// 	return nil, nil
	// }

	return nil, nil
}
