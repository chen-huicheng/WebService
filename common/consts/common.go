package consts

import "time"

// 登录相关 token cookie session
const (
	CookieName   = "chen_session"
	CookieStrLen = 25
	CookieMaxAge = 24 * time.Hour
	UserKey      = "user"
)
