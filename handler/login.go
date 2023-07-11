package handler

import (
	"WebService/dal/dto"
	"WebService/service"

	"WebService/common/response"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		response.Response(c, err, nil)
		return
	}
	ss := service.GetService()
	resp, err := ss.Login(c, &req)
	if err != nil {
		response.Response(c, err, nil)
		return
	}
	response.Response(c, nil, resp)
}
