package service

import (
	"errors"
	"rest-api/entity"
	"rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	AddPost(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	DeleteOne(id int64) error
	GetOne(id int64) (entity.Post, error)
	UpdatePost(post *entity.Post) (*entity.Post, error)
}

type service struct{}

var (
	repos repository.PostRepo
)

func NewPostService(repo repository.PostRepo) PostService {
	repos = repo
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The title is empty")
		return err
	}
	return nil
}

func (*service) AddPost(post *entity.Post) (*entity.Post, error) {
	// post.ID = rand.Int63()
	return repos.Save(post)

}

func (*service) FindAll() ([]entity.Post, error) {
	return repos.FindAll()

}

func (*service) DeleteOne(id int64) error {
	return repos.DeleteOne(id)

}

func (*service) GetOne(id int64) (entity.Post, error) {
	return repos.GetOne(id)
}

func (*service) UpdatePost(post *entity.Post) (*entity.Post, error) {
	return repos.Update(post)
}
