package common

import (
	"fmt"

	"fanfan.me/DoubanSpider/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	// args := fmt.Sprintf("%s:%s@tcp(%s:%s)%s?charset=%s&parseTime=true",
	// 	username,
	// 	password,
	// 	host,
	// 	port,
	// 	database,
	// 	charset,
	// )
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset) //格式化输出
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect databse,err:" + err.Error())
	}

	db.AutoMigrate(&model.Movie{})
	DB = db
	return DB
}
func GetDb() *gorm.DB {
	return DB
}
