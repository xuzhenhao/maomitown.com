package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"maomitown.com/handler"
	"maomitown.com/model"
	"maomitown.com/pkg/errno"
)

// Update 更新用户记录
func Update(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("id"))
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ParamBindError, nil)
		return
	}

	u.ID = uint64(userID)
	//校验数据格式
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ValidationError, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.EncryptPwdError, nil)
		return
	}

	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.DatabaseError, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
