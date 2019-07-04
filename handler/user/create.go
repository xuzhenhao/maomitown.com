package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	"maomitown.com/handler"
	"maomitown.com/pkg/errno"
)

// Create 创建一个新的账户
func Create(c *gin.Context) {
	var r CreateRequest

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ParamBindError})
		handler.SendResponse(c, errno.ParamBindError, nil)
		return
	}

	//解析路由中的参数
	usrname := c.Param("username")
	log.Infof("URL 中的用户名为: %s", usrname)

	//解析路由中的querystring
	desc := c.Query("desc")
	log.Infof("URL 中query参数desc的值为:%s", desc)

	//获取请求头中的值
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Contenttype: %s", contentType)

	log.Debugf("将要创建的用户名:%s 密码:%s", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.UserNotFoundError, fmt.Errorf("用户名不能为空"))
		handler.SendResponse(c, err, nil)
		return
	}
	if r.Password == "" {
		err = fmt.Errorf("密码为空")
		handler.SendResponse(c, err, nil)
	}

	//构建返回的数据
	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, err, rsp)
}
