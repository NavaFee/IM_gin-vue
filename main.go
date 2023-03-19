package main

import (
	"github/IM_gin+vue/router"
	"github/IM_gin+vue/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	// fmt.Println(viper.GetString("mysql.dns"))
	// fmt.Println(viper.GetString("redis.addr"))
	// fmt.Println(viper.GetInt("port.udp"))
	// fmt.Println(viper.GetInt("port.udp"))
	r := router.Router()
	r.Run(":8080")
}
