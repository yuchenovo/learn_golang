package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

func main() {
	ReadFromConfig()
}

func ReadFromConfig() {
	config := viper.New()            //创建一个viper文件
	config.AddConfigPath("./config") //配置文件所在目录
	config.SetConfigName("account")  //文件名
	config.SetConfigType("yaml")     //文件类型，配置文件的后缀
	//监控配置文件变化
	config.WatchConfig()
	config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
	})
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("找不到配置文件")
		} else {
			panic(err)
		}
	}
	for {
		PrintConfig(config)
		time.Sleep(3 * time.Second)
	}

}
func PrintConfig(config *viper.Viper) {
	//读取配置
	user1 := config.GetString("section1.user")
	user2 := config.GetString("section2.user")
	height := config.GetInt32("section1.body.height")
	weight := config.GetInt32("section1.body.weight")
	fmt.Println(user1, user2, height, weight)
}
