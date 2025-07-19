package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局配置变量
var Conf = new(Config)

type Config struct {
	App   *AppConfig   `mapstructure:"app"`   // 对应 YAML 中的 "app"
	Log   *LogConfig   `mapstructure:"log"`   // 对应 YAML 中的 "log"
	MySQL *MySQLConfig `mapstructure:"mysql"` // 对应 YAML 中的 "mysql"
	Redis *RedisConfig `mapstructure:"redis"` // 对应 YAML 中的 "redis"
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`       //日志级别
	Filename   string `mapstructure:"filename"`    //日志文件名
	MaxSize    int    `mapstructure:"max_size"`    //日志文件最大尺寸，单位MB
	MaxBackups int    `mapstructure:"max_backups"` //日志文件最大备份数
	MaxAge     int    `mapstructure:"max_age"`     //日志文件最大保存天数
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`           // MySQL服务器地址
	Port         int    `mapstructure:"port"`           // MySQL服务器端口
	User         string `mapstructure:"user"`           // MySQL用户名
	Password     string `mapstructure:"password"`       // MySQL密码
	DBName       string `mapstructure:"dbname"`         // MySQL数据库名
	MaxOpenConns int    `mapstructure:"max_open_conns"` // MySQL最大打开连接数
	MaxIdleConns int    `mapstructure:"max_idle_conns"` // MySQL最大空闲连接数
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`      // Redis服务器地址
	Port     int    `mapstructure:"port"`      // Redis服务器端口
	Password string `mapstructure:"password"`  // Redis密码
	DB       int    `mapstructure:"db"`        // Redis数据库编号
	PoolSize int    `mapstructure:"pool_size"` // Redis连接池大小
}

func Init() (err error) {
	viper.SetConfigName("config") //指定配置文件名
	viper.SetConfigType("yaml")   //指定配置文件类型
	viper.AddConfigPath(".")      //指定配置文件路径
	err = viper.ReadInConfig()    //读取配置文件
	if err != nil {
		// 如果配置文件读取失败，打印错误信息并退出
		fmt.Printf("Fatal error config file: %s \n", err)
		panic("Fatal error config file: " + err.Error())
	}
	// 将配置文件内容反序列化到Conf结构体中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("Unmarshal config failed, err:%v\n", err)
	}
	viper.WatchConfig() // 监听配置文件变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("Unmarshal config failed, err:%v\n", err)
		} else {
			fmt.Println("Config reloaded successfully")
		}
	})
	return
}
