package store

import (
	"context"
	"database/sql"
)

type Follower struct {
	FollowerID int64  `json:"follower_id"`
	UserID     int64  `json:"user_id"`
	CreatedAt  string `json:"created_at"`
}

type FollowerStorage struct {
	db *sql.DB
}

func (s *FollowerStorage) Follow(ctx context.Context, followerID, userID int64) error {
	query := `
		insert into followers (follower_id, user_id)
		values ($1, $2)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}

func (s *FollowerStorage) Unfollow(ctx context.Context, followerID, userID int64) error {
	query := `
		delete from followers
		where follower_id = $1 and user_id = $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}
