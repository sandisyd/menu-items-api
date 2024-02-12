package main

import (
	"menu/menuItems/config"
	menuitemscontroller "menu/menuItems/menu_items_controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDb()
	r.GET("api/getDataMenu", menuitemscontroller.GetDataMenu)
	r.Run(":3004")
}
