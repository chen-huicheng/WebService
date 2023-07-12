package session

import (
	"WebService/common/consts"
	"WebService/dal/do"

	"github.com/gin-gonic/gin"
)

// GetCurUser 从 gin.context 中获取用户信息
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

// SetCurUser 用户信息添加到 gin.context 中
func SetCurUser(c *gin.Context, user *do.User) {
	c.Set(consts.UserKey, user)
}
