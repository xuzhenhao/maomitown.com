package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"maomitown.com/handler"
	"maomitown.com/model"
	"maomitown.com/pkg/errno"
)

// Delete 删除用户
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userID)); err != nil {
		handler.SendResponse(c, errno.DatabaseError, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
