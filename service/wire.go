//go:build wireinject
// +build wireinject

package service

import (
	"WebService/dal/dao"
	"WebService/dal/dao_impl"

	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	wire.Struct(new(CommonService), "*"),
)

var daoSet = wire.NewSet(
	wire.Struct(new(dao_impl.UserDaoImpl), "*"),

	wire.Bind(new(dao.UserDao), new(*dao_impl.UserDaoImpl)),
)

// var driverSet = wire.NewSet(
// 	rdb.InitRedis,
// )

type Component struct {
	CommonService
}

func InitService() (*Component, error) {
	panic(wire.Build(
		serviceSet,
		daoSet,
		wire.Struct(new(Component), "*"),
	))
}
