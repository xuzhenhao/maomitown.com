package model

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

// Database 数据库
type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

// DB 数据库实例
var DB *Database

// 连接数据库
func openDB(username, password, addr, name string) *gorm.DB {
	//拼接连接数据库的字符串
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "数据库连接失败，数据库名:", name)
	}
	setupDB(db)

	return db
}

// 设置数据库
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//设置最大打开的连接数，可以避免并发太高导致连接错误
	db.DB().SetMaxOpenConns(20000)
	//设置闲置的连接数。当开启的一个连接使用完成后可以放在池里等候下一次使用
	db.DB().SetMaxIdleConns(0)
}

// GetSelfDB 连接本地数据库
func GetSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

// GetDockerDB 获取容器内数据库
func GetDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

// Init 初始化方法
func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

// Close 关闭数据库
func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}
