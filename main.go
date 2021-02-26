package main

import (
	"os"

	"fanfan.me/DoubanSpider/common"
	"fanfan.me/DoubanSpider/spider"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	spider.Spider(db)
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
