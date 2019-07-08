package model

import (
	validator "gopkg.in/go-playground/validator.v9"
	"maomitown.com/pkg/auth"
	"maomitown.com/pkg/constvar"
)

/* 用户表模型定义及增删改查 */

// UserModel 用户模型
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// TableName user表名
func (u *UserModel) TableName() string {
	return "tb_users"
}

// Create 往数据库中插入一个新的用户记录
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser 根据用户id删除用户记录
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.ID = id
	return DB.Self.Delete(&user).Error
}

// Update 更新用户记录
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser 根据用户名查询获取用户信息
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)

	return u, d.Error
}

// ListUser 查询用户列表
func ListUser(offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64
	//查询个数
	if err := DB.Self.Model(&UserModel{}).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := DB.Self.Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil

}

// Validate 校验UserModel数据格式
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// Encrypt 加密用户密码
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
