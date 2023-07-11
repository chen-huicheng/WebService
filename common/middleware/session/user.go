package session

import (
	"WebService/common/consts"
	"WebService/dal/do"

	"github.com/gin-gonic/gin"
)

func GetCurUser(c *gin.Context) *do.User {
	val, exists := c.Get(consts.UserKey)
	if !exists {
		return nil
	}
	if v, ok := val.(*do.User); ok {
		return v
	}
	return nil
}
func SetCurUser(c *gin.Context, user *do.User) {
	c.Set(consts.UserKey, user)
}
