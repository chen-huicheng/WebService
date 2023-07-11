package middleware

import (
	"WebService/common"
	"WebService/common/consts"
	"WebService/common/middleware/session"
	"WebService/common/response"
	"WebService/dal/dao"
	"WebService/dal/driver/rdb"

	"github.com/gin-gonic/gin"
)

// 登录授权中间件
/*
	从 cookie 中拿 token，然后从 Redis 中获取 userID
	然后查询 user 并将 user 信息存到 gin.Context 中方便后续使用
*/
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullPath := c.FullPath()
		if isNoNeedAuth(fullPath) {
			c.Next()
			return
		}
		cookie, err := c.Request.Cookie(consts.CookieName)
		if err != nil {
			response.Response(c, common.BizErr.SetMsg("未登录"), nil)
			return
		}
		token := cookie.Value
		ctx := c.Request.Context()
		userID, err := rdb.Get(ctx, token)
		if err != nil || userID == "" {
			response.Response(c, common.BizErr.SetMsg("未登录"), nil)
			return
		}
		uDao := dao.NewUserDao(ctx)
		user, err := uDao.GetUserByID(ctx, userID)
		if err != nil {
			response.Response(c, common.BizErr.SetMsg("未登录"), nil)
			return
		}
		session.SetCurUser(c, user)
	}
}

func isNoNeedAuth(path string) bool {
	return path == "/user/login"

}
