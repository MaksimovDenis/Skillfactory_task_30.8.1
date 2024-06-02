package postsPG

import (
	"context"
	"errors"

	"github.com/MaksimovDenis/skillfactory_task_30.8.1/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostPostgres struct {
	db *pgxpool.Pool
}

func NewPostPostgres(db *pgxpool.Pool) *PostPostgres {
	return &PostPostgres{db: db}
}

func (p *PostPostgres) Posts() ([]models.Post, error) {
	rows, err := p.db.Query(context.Background(), `
		SELECT 
			id,
			author_id,
			title,
			content, 
			created_at
		FROM posts
		ORDER BY id;
	`)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		err = rows.Scan(
			&p.ID,
			&p.AuthorID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (p *PostPostgres) AddPost(post models.Post) error {

	_, err := p.db.Exec(context.Background(), `
	INSERT INTO posts (author_id, title, content)
	VALUES ($1, $2, $3);
	`, post.AuthorID, post.Title, post.Content)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return errors.New("invalid author_id: doesn't exist")
		}
		return err
	}

	return nil
}

func (p *PostPostgres) UpdatePost(post models.UpdatePost) error {
	_, err := p.db.Exec(context.Background(), `
	UPDATE posts
	SET author_id = $1, title = $2, content = $3
	WHERE id = $4;
	`, post.AuthorID, post.Title, post.Content, post.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return errors.New("failed to update post")
		}
		return err
	}

	return nil
}

func (p *PostPostgres) DeletePost(id int) error {
	cmdTag, err := p.db.Exec(context.Background(), `
	DELETE FROM posts 
	WHERE id=$1
	`,
		id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() < 1 {
		return errors.New("failed to delete post")
	}

	return nil
}
