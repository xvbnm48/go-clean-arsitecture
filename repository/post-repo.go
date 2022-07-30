package repository

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/xvbnm48/go-clean-arsitecture/entity"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]*entity.Post, error)
}

type repo struct {
}

// NewRepository

func NewRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "golang-clean-arsitecture"
	CollectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/fariz/FIREBASE/golang-clean-arsitecture-firebase-adminsdk-m2cjx-ca1b51b1e7.json")
	client, err := firestore.NewClient(ctx, "golang-clean-arsitecture")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	// defer client.Close()

	_, _, err = client.Collection(CollectionName).Add(ctx, map[string]interface{}{
		"id":    post.Id,
		"title": post.Title,
		"text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed to add new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]*entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []*entity.Post
	itr := client.Collection(CollectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		post := entity.Post{
			Id:    doc.Data()["id"].(int64),
			Title: doc.Data()["title"].(string),
			Text:  doc.Data()["text"].(string),
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
