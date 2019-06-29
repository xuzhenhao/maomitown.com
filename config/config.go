package config

import (
	"strings"

	"github.com/lexkong/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 配置项
type Config struct {
	Name string
}

// Init 初始化配置
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	if err := c.initConfig(); err != nil {
		return err
	}

	//初始化日志配置
	c.initLog()
	//监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		//指定了配置文件，解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		//没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	//设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	//读取匹配的环境变量
	viper.AutomaticEnv()
	//读取环境变量的前缀
	viper.SetEnvPrefix("MMTSERVER")
	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		//viper 解析配置文件
		return err
	}

	return nil
}

// watchConfig 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

// initLog 初始化日志配置
func (c *Config) initLog() {
	logCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_level"),
		LogFormatText:  viper.GetBool("log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&logCfg)
}
