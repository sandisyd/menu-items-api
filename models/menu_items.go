package models

import "github.com/jinzhu/gorm"

type MenuItems struct {
	gorm.Model
	Name    string    `gorm:"not null;" json:"name"`
	Icon    string    `gorm:"type:text" json:"icon"`
	Submenu []Submenu `gorm:"foreignKey:MenuItemId" json:"submenu"`
}

type Submenu struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `gorm:"type:char(255)" json:"description"`
	Icon        string `gorm:"type:text" json:"icon"`
	ActivityApp string `json:"activity_app"`
	RouteWeb    string `json:"route_web"`
	Publish     bool   `json:"publish"`
	Asn         bool   `json:"asn"`
	Nik         bool   `json:"nik"`
	MenuItemId  uint   `json:"menu_item_id"`
}
