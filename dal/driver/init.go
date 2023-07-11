package driver

import "WebService/dal/driver/rdb"

func Init() {
	rdb.InitRedis()
}
