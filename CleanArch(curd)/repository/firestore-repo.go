package repository

// import (
// 	"context"
// 	"log"
// 	"rest-api/entity"

// 	"cloud.google.com/go/firestore"
// )

// // to implement the interface create an struct
// type firerepo struct{}

// // constructor to create a new instance
// // change func name to firestore
// func NewFireStoreRepo() PostRepo {
// 	return &firerepo{}
// }

// // id and name from firestore database
// const (
// 	projectId      string = "pragmatic-reviewss"
// 	collectionName string = "posts"
// )

// func (*firerepo) Save(post *entity.Post) (*entity.Post, error) {
// 	//returns an ampty context
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, projectId)
// 	if err != nil {
// 		log.Fatal("Failed to create a Firestore client", err)
// 		return nil, err
// 	}

// 	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
// 		"ID":    post.ID,
// 		"Title": post.Title,
// 		"Text":  post.Text,
// 	})
// 	if err != nil {
// 		log.Fatal("Failed in adding a new post", err)
// 		return nil, err
// 	}
// 	defer client.Close()

// 	return post, nil

// }

// func (*firerepo) FindAll() ([]entity.Post, error) {

// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, projectId)
// 	if err != nil {
// 		log.Fatal("Failed to create a Firestore client", err)
// 		return nil, err
// 	}
// 	defer client.Close()
// 	var posts []entity.Post
// 	//itterate through every post
// 	iterater := client.Collection(collectionName).Documents(ctx)
// 	for {
// 		doc, err := iterater.Next()
// 		if err != nil {
// 			log.Fatal("Failed to iterate through posts", err)
// 			return nil, err
// 		}
// 		post := entity.Post{
// 			ID:    doc.Data()["ID"].(int64),
// 			Title: doc.Data()["Title"].(string),
// 			Text:  doc.Data()["Text"].(string),
// 		}
// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }
