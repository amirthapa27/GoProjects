package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Items struct {
	ItemID         uint64 `json:"item_id" gorm:"primaryKey"`
	MenuCategoryID uint64 `json:"menu_category_id" gorm:"refrences:category_id"`
	Name           string `json:"name"`
	// Photo
	Price       float64 `json:"price"`
	Ingredients string  `json:"ingredients"`
}
type Menu struct {
	CategoryID uint64 `json:"category_id" gorm:"primaryKey"`
	Category   string `json:"category"`
	// Items Items `json:"items"`
}

func Setup() {
	dsn := "host=172.17.0.2 user=postgres password=mypass dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Menu{}, &Items{})
	if err != nil {
		panic(err)
	}
}

// getting all the categories
func GetMenu() ([]Menu, error) {
	var menu []Menu
	tx := db.Find(&menu)
	if tx.Error != nil {
		return []Menu{}, tx.Error
	}
	return menu, nil
}

// creating new category
func CreateCategory(menu Menu) error {
	tx := db.Create(&menu)
	return tx.Error
}

// delete a category
func DeleteCategory(id uint64) error {
	tx := db.Unscoped().Delete(&Menu{}, id)
	return tx.Error
}

// update category
func UpdateMenu(menu Menu) error {
	tx := db.Save(&menu)
	return tx.Error
}

// getting all items
func GetAllItems() ([]Items, error) {
	var items []Items
	tx := db.Find(&items)
	if tx.Error != nil {
		return []Items{}, tx.Error
	}
	return items, nil
}

// creating an item
func CreateItem(item Items) error {
	tx := db.Create(&item)
	return tx.Error
}

// update items
func UpdateItems(item Items) error {
	tx := db.Save(&item)
	return tx.Error
}

// delete item
func DeleteItem(id uint64) error {
	tx := db.Unscoped().Delete(&Items{}, id)
	return tx.Error
}

// get category
func GetCategory(id uint64) ([]Items, error) {
	var items []Items
	tx := db.Joins("inner join menus on menus.category_id=items.menu_category_id").Where("items.menu_category_id = ?", id).Find(&items)

	if tx.Error != nil {
		return []Items{}, tx.Error
	}
	return items, nil
}
