package repository

import (
	"rest-api/entity"
)

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	GetOne(id int64) (entity.Post, error)
	DeleteOne(id int64) error
	Update(post *entity.Post) (*entity.Post, error)
}
