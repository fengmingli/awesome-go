/**
 * @Author LFM
 * @Date 2021/10/21 4:29 下午
 * @Since V1
 */

package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

func VipReadConfig() {
	//文件名
	viper.SetConfigName("config")
	//文件类型
	viper.SetConfigType("yaml")
	//配置文件所在路径
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Viper read config file: %s \n", err))
	}
	config := viper.GetString("awesome.test")
	fmt.Println(config)
}
