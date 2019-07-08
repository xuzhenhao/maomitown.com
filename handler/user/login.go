package user

import (
	"github.com/gin-gonic/gin"
	"maomitown.com/handler"
	"maomitown.com/model"
	"maomitown.com/pkg/auth"
	"maomitown.com/pkg/errno"
	"maomitown.com/pkg/token"
)

// Login 用户登录接口
func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ParamBindError, nil)
		return
	}

	//查询用户信息
	d, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.UserNotFoundError, nil)
		return
	}

	//验证密码是否正确
	if err := auth.Compare(d.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.PwdIncorrectError, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.ID, Username: d.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.TokenError, nil)
		return
	}

	handler.SendResponse(c, nil, model.Token{Token: t})
}
