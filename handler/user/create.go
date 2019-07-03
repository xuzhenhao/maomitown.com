package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"maomitown.com/pkg/errno"
)

// Create 创建一个新的账户
func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ParamBindError})
		return
	}

	log.Debugf("将要创建的用户名:%s 密码:%s", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.UserNotFoundError, fmt.Errorf("用户名不能为空"))
		log.Errorf(err, "格式错误")
	}

	if code, _ := errno.DecodeErr(err); code == errno.UserNotFoundError.Code {
		log.Debug("用户没找到错误")
	}

	if r.Password == "" {
		err = fmt.Errorf("密码为空")
	}
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
