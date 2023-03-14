package main

import (
	"github/IM_gin+vue/router"
	"github/IM_gin+vue/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8080")
}
