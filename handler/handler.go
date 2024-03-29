package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"maomitown.com/pkg/errno"
)

// Response 定义返回格式，Code和Message通过DecodeErr解析err获得，Data可以根据业务自行定义(map,array等)
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse 全局统一的返回函数，格式化返回
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

/*
test:
 curl -XPOST -H "Content-Type:application/json" http://127.0.0.1:8080/v1/user -d'{"username":"admin","password":"123456"}'
 curl -XDELETE -H "Content-Type:application/json" http://127.0.0.1:8080/v1/user/2
 curl -XPUT -H "Content-Type:application/json" http://127.0.0.1:8080/v1/user/2 -d'{"username":"cloudmodify","password":"123456"}'
 curl -XGET -H "Content-Type:application/json" http://127.0.0.1:8080/v1/user/list

 curl -XPOST -H "Content-Type:application/json" http://127.0.0.1:8080/login -d'{"username":"modify","password":"123456"}'

 curl -XGET -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NjI1NzM0OTUsImlkIjozLCJuYmYiOjE1NjI1NzM0OTUsInVzZXJuYW1lIjoibW9kaWZ5In0.zKOhCh0mPCCOafxDG5qtkf5n4jwUNAeAjhhx3x609MQ" -H "Content-Type:application/json" http://127.0.0.1:8080/v1/user/list
*/
