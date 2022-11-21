package repository

import (
	"rest-api/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type psqlrepo struct{}

func NewPSQLRepo() PostRepo {
	dsn := "host=172.17.0.2 user=postgres password=mypass dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(entity.Post{})
	if err != nil {
		panic(err)
	}
	return &psqlrepo{}
}

// func Setup() {
// 	dsn := "host=172.17.0.2 user=postgres password=mypass dbname=postgres port=5432 sslmode=disable"
// 	var err error
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = db.AutoMigrate(entity.Post{})
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (*psqlrepo) Save(post *entity.Post) (*entity.Post, error) {
	tx := db.Create(&post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}

func (*psqlrepo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	tx := db.Find(&posts)
	if tx.Error != nil {
		return []entity.Post{}, tx.Error
	}
	return posts, tx.Error
}

func (*psqlrepo) DeleteOne(id int64) error {
	tx := db.Unscoped().Delete(&entity.Post{}, id)
	return tx.Error
}

func (*psqlrepo) Update(post *entity.Post) (*entity.Post, error) {
	tx := db.Save(&post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}

func (*psqlrepo) GetOne(id int64) (entity.Post, error) {
	var post entity.Post
	tx := db.Where("id = ?", id).Find(&post)
	if tx.Error != nil {
		return entity.Post{}, tx.Error
	}
	return post, nil
}
