package main

import (
	"fmt"
	"github/IM_gin+vue/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("config app inited ...")

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//迁移schema
	db.AutoMigrate(&models.UserBasic{})

	//Create
	user := &models.UserBasic{}
	user.Name = "贺飞"
	// user.LoginTime = time.Now()
	// user.HeartbeatTime = time.Now()

	db.Create(user)

	// //Read
	// fmt.Println(db.First(user, 1))

	// //update
	// db.Model(user).Update("PassWord", "1234")
	// fmt.Println("config app:", viper.Get("app"))
	// fmt.Println("config mysql:", viper.Get("mysql"))

}
