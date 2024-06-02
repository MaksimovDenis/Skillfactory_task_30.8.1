package postsMongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"skillfactory_task_31.3.1/internal/models"
)

const (
	databaseName   = "skillfactory" // имя учебной БД
	collectionName = "posts"        // имя коллекции в учебной БД
)

type PostMongo struct {
	db *mongo.Client
}

func NewPostMongo(db *mongo.Client) *PostMongo {
	return &PostMongo{db: db}
}

func (p *PostMongo) Posts() ([]models.Post, error) {
	collection := p.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var posts []models.Post

	for cur.Next(context.Background()) {
		var p models.Post
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, cur.Err()
}

func (p *PostMongo) AddPost(post models.Post) error {
	collection := p.db.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostMongo) UpdatePost(post models.UpdatePost) error {
	collection := p.db.Database(databaseName).Collection(collectionName)

	filter := bson.M{"id": post.ID}

	update := bson.M{"$set": bson.M{"title": post.Title, "content": post.Content, "author_id": post.AuthorID}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostMongo) DeletePost(id int) error {
	collection := p.db.Database(databaseName).Collection(collectionName)

	filter := bson.M{"id": id}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("no document found with id: %v", id)
	}

	return nil
}
