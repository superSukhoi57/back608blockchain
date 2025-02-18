package firstInit

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func init() {

	log.Println("viper开始初始化！")
	//TODO：路径还是在go.mod那里开始算！
	viper.SetConfigFile("./conf.yml") // 指定配置文件路径
	viper.SetConfigName("conf")       // 配置文件名称(无扩展名)
	viper.SetConfigType("yml")        // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")          // 指定在那里照配置文件，可以写多个路径。这里“.”代表在工作目录中查找配置
	err := viper.ReadInConfig()       // 查找并读取配置文件
	if err != nil {                   // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
