package menuitemscontroller

import (
	"fmt"
	"menu/menuItems/config"
	"menu/menuItems/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
)

func GetDataMenu(c *gin.Context) {
	var menuItems []models.MenuItems
	config.DB.Debug().Preload("Submenu").Find(&menuItems)
	// config.DB.Preload("Submenu").Find(&menuItems)
	fmt.Println("items menu : ", menuItems)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "get data success",
		"data":    menuItems})
}
