package service

import (
	"WebService/common"
	"WebService/common/consts"
	"WebService/common/middleware/session"
	"WebService/common/utils"
	"WebService/dal/dao"
	"WebService/dal/driver/rdb"
	"WebService/dal/dto"

	"github.com/gin-gonic/gin"
)

type CommonService struct {
	userDao dao.UserDao
}

func (s *CommonService) Login(c *gin.Context, req *dto.LoginReq) (*dto.LoginResp, error) {
	ctx := c.Request.Context()
	user, err := s.userDao.GetUserByID(ctx, req.User)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, common.BizErr.SetMsg("用户不存在")
	}
	if user.Passwd != req.Passwd {
		return nil, common.BizErr.SetMsg("密码错误")
	}
	token := utils.RandString(consts.CookieStrLen)
	rdb.Set(ctx, token, user.ID, consts.CookieMaxAge)
	c.SetCookie(consts.CookieName, token, int(consts.CookieMaxAge), "/user", "localhost", false, true)
	return &dto.LoginResp{}, nil
}

func (s *CommonService) ReSetPasswd(c *gin.Context, req *dto.ReSetPasswdReq) (*dto.ReSetPasswdResp, error) {
	ctx := c.Request.Context()
	user := session.GetCurUser(c)
	if user == nil {
		return nil, common.BizErr.SetMsg("未登录")
	}
	if user.Passwd != req.OldPasswd {
		return nil, common.BizErr.SetMsg("密码错误")
	}
	user.Passwd = req.NewPasswd
	err := s.userDao.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	c.SetCookie(consts.CookieName, "", -1, "/user", "localhost", false, true)
	return &dto.ReSetPasswdResp{}, nil
}
