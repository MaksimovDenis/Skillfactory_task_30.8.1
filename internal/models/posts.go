package models

import "time"

type Author struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Post struct {
	ID        int       `json:"id" db:"id" bson:"id,omitempty"`
	Title     string    `json:"title" db:"title" bson:"title"`
	Content   string    `json:"content" db:"content" bson:"content"`
	AuthorID  int       `json:"author_id" db:"author_id" bson:"author_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at" bson:"created_at"`
}

type UpdatePost struct {
	ID       int    `json:"id" db:"id" bson:"id,omitempty"`
	Title    string `json:"title" db:"title" bson:"title"`
	Content  string `json:"content" db:"content" bson:"content"`
	AuthorID int    `json:"author_id" db:"author_id" bson:"author_id"`
}
