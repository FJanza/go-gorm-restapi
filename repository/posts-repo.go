package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/fjanza/go-gorm-restapi/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Find(id int) (entity.Post, error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "janza-reviews"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile("D:/repositorios/go-gorm-restapi/janza-reviews-firebase.json"))
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	defer client.Close()

	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	// Add the new post to Firestore
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile("D:/repositorios/go-gorm-restapi/janza-reviews-firebase.json"))

	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post

	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (*repo) Find(id int) (entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile("D:/repositorios/go-gorm-restapi/janza-reviews-firebase.json"))

	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return entity.Post{}, err
	}

	defer client.Close()

	var post entity.Post

	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return entity.Post{}, err
		}

		if doc.Data()["ID"].(int64) == int64(id) {
			post = entity.Post{
				ID:    doc.Data()["ID"].(int64),
				Title: doc.Data()["Title"].(string),
				Text:  doc.Data()["Text"].(string),
			}
		}
	}
	return post, nil
}
