package service

import (
	"fmt"
	"sync"
)

var serviceOnce sync.Once
var componentService *Component

func GetService() *Component {
	if componentService == nil {
		componentService = Init()
	}
	return componentService
}
func Init() *Component {
	serviceOnce.Do(func() {
		var err error
		componentService, err = InitService()
		if err != nil {
			panic(fmt.Sprintf("init service err:%v", err))
		}
	})
	return componentService
}
