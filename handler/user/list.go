package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"maomitown.com/handler"
	"maomitown.com/service/userservice"
)

type ListResponse struct {
	TotalCount uint64                  `json:"totalCount"`
	UserList   []*userservice.UserInfo `json:"userList"`
}

// List 查询用户列表
func List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	infos, count, err := userservice.ListUser(offset, limit)

	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
