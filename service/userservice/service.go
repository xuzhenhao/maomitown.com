package userservice

import (
	"fmt"
	"sync"

	"maomitown.com/model"
)

// UserInfo 用户信息
type UserInfo struct {
	ID       uint64 `json:"id"`
	UserName string `json:"username"`
	SayHello string `json:"sayHello"`
}

// UserList 用户信息列表
type UserList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*UserInfo
}

// ListUser 获取用户列表
func ListUser(offset, limit int) ([]*UserInfo, uint64, error) {
	infos := make([]*UserInfo, 0)
	users, count, err := model.ListUser(offset, limit)

	// log.Infof("数据库查询到的用户列表 %v", users)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := UserList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*UserInfo, len(users)),
	}

	// errChan := make(chan error, 1)
	finishChan := make(chan bool, 1)

	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IDMap[u.ID] = &UserInfo{
				ID:       u.ID,
				UserName: u.Username,
				SayHello: fmt.Sprintf("hello,%s", u.Username),
			}
			// log.Infof("用户 %v", userList.IDMap[u.ID])
		}(u)
	}

	wg.Wait()
	close(finishChan)

	for _, id := range ids {
		infos = append(infos, userList.IDMap[id])
	}

	// log.Infof("返回结果 %v", infos)
	return infos, count, nil

}
