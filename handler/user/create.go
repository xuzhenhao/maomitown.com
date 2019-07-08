package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	"maomitown.com/handler"
	"maomitown.com/model"
	"maomitown.com/pkg/errno"
)

// Create 创建一个新的账户
func Create(c *gin.Context) {

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ParamBindError, nil)
		return
	}

	log.Debugf("将要创建的用户名:%s 密码:%s", r.Username, r.Password)

	//数据填充模型
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	//校验数据
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ValidationError, nil)
		return
	}

	//加密用户密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.EncryptPwdError, nil)
		return
	}

	//插入用户数据到数据库
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.DatabaseError, nil)
		return
	}

	//构建返回的数据
	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, rsp)
}
