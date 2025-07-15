package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

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
	viper.WatchConfig() // 监听配置文件变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:")
	})
	return
}
