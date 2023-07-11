package handler

import (
	"WebService/common/response"
	"WebService/dal/dto"
	"WebService/service"

	"github.com/gin-gonic/gin"
)

func ResetPasswd(c *gin.Context) {
	var req dto.ReSetPasswdReq
	if err := c.ShouldBind(&req); err != nil {
		response.Response(c, err, nil)
		return
	}
	ss := service.GetService()
	resp, err := ss.ReSetPasswd(c, &req)
	if err != nil {
		response.Response(c, err, nil)
		return
	}
	response.Response(c, nil, resp)
}
