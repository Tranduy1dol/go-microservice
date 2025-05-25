package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int64     `json:"user_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Tags      []string  `json:"tags"`
	Comments  []Comment `json:"comments"`
}

type PostStorage struct {
	db *sql.DB
}

func (s *PostStorage) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (title, content, user_id, tags) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at;
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.UserID,
		pq.Array(&post.Tags),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStorage) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT id, title, content, user_id, created_at, updated_at, tags 
		FROM posts WHERE id = $1;
	`

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStorage) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostStorage) Update(ctx context.Context, post *Post) error {
	query := `
		UPDATE posts 
		SET title = $1, content = $2
		where id = $3
	`

	_, err := s.db.ExecContext(ctx, query, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}

	return nil
}
