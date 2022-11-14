package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Stocks struct {
	gorm.Model
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}

func Setup() {
	dsn := "host=172.17.0.2 user=admin password=mypass dbname=admin port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Stocks{})
	if err != nil {
		panic(err)
	}

}

func GetAllStocks() ([]Stocks, error) {
	var stocks []Stocks
	tx := db.Find(&stocks)
	if tx.Error != nil {
		return []Stocks{}, tx.Error
	}
	return stocks, nil
}

func GetOneStock(id uint64) (Stocks, error) {
	var stock Stocks
	tx := db.Where("ID = ?", id).First(&stock)
	if tx.Error != nil {
		return Stocks{}, tx.Error
	}
	return stock, nil
}

func CreateStock(stock Stocks) error {
	tx := db.Create(&stock)
	return tx.Error
}

func DeleteStock(id uint64) error {
	// var stock Stocks
	// tx := db.Where("ID = ?", id).Delete(&stock)
	tx := db.Unscoped().Delete(&Stocks{}, id)
	return tx.Error
}

func UpdateStocks(stock Stocks) error {
	tx := db.Save(&stock)
	return tx.Error
}
