package response

import (
	"WebService/common"
	"log"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Response 统一输出响应
func Response(c *gin.Context, err error, resp interface{}) {

	httpCode := 200
	var statusCode int
	var message string
	if err == nil {
		statusCode = 200
		message = "success"
	} else if e, ok := err.(common.HTTPErr); ok {
		statusCode = e.Code()
		message = e.Msg()
		log.Printf("[Error] path:%s err:%v", c.FullPath(), e)
	} else {
		statusCode = 500
		message = "服务错误"
		log.Printf("[Error] path:%s err:%v", c.FullPath(), err)
	}
	c.PureJSON(httpCode, Result{
		Code: statusCode,
		Data: resp,
		Msg:  message,
	})
}
