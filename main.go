package main

import (
	"WebService/conf"
	"WebService/dal/driver"

	"github.com/gin-gonic/gin"
)

func main() {
	driver.Init()
	r := gin.Default()
	register(r)
	cf := conf.GetConfig()
	addr := cf.Server.Host + ":" + cf.Server.Port
	r.Run(addr) // listen and serve on 127.0.0.1:8080
}
