package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"maomitown.com/config"
	"maomitown.com/model"
	"maomitown.com/router"
	"maomitown.com/router/middleware"
)

var (
	// cfg变量值从命令行flag传入，也可以为空，如果为空，默认读取conf/config.yaml
	cfg = pflag.StringP("config", "c", "", "server config file path")
)

func main() {
	// 项目配置初始化
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("runmode"))

	//创建web服务器
	g := gin.New()

	//初始化路由配置
	router.Load(
		g,
		middleware.RequestID(),
	)
	//另起一个协程检查服务器是否能正常访问
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("服务无响应", err)
		}
		log.Info("路由部署成功")
	}()

	log.Infof("开始监听位于 %s 端口的http请求", viper.GetString("addr"))
	http.ListenAndServe(viper.GetString("addr"), g)

}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Infof("%v", err)
		//休眠后重试
		time.Sleep(time.Second)
	}
	return errors.New("无法连接到路由")
}
