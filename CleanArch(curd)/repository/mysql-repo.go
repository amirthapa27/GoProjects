package repository

// import (
// 	"rest-api/entity"

// 	"gorm.io/gorm"

// 	"gorm.io/driver/mysql"
// )

// var (
// // db *gorm.DB
// )

// type mysqlrepo struct{}

// func NewmysqlRepo() PostRepo {
// 	var err error
// 	dsn := ("root:mypass@/mysql?charset=utf8mb4&parseTime=True&loc=Local")
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = db.AutoMigrate(&entity.Post{})
// 	if err != nil {
// 		panic(err)

// 	}

// 	return &mysqlrepo{}
// }

// func (*mysqlrepo) Save(post *entity.Post) (*entity.Post, error) {
// 	tx := db.Create(&post)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return post, nil
// }

// func (*mysqlrepo) FindAll() ([]entity.Post, error) {
// 	var posts []entity.Post
// 	tx := db.Find(&posts)
// 	if tx.Error != nil {
// 		return []entity.Post{}, tx.Error
// 	}
// 	return posts, nil
// }
