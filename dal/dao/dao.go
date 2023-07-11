package dao

import (
	"WebService/dal/dao_impl"
	"context"
)

func NewUserDao(ctx context.Context) UserDao {
	return &dao_impl.UserDaoImpl{}
}
