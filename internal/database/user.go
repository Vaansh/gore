package database

import (
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	r := &UserRepository{db: db}
	if err := r.createUser(model.NewUser("", "")); err != nil {
		return nil
	}
	return r
}

func (r *UserRepository) createUser(user *model.User) error {
	tableName := getTableName(user)
	query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
			post_id VARCHAR(20) NOT NULL,
			source_id VARCHAR(50) NOT NULL,
			author_name VARCHAR(100),
            platform VARCHAR(5) NOT NULL,
            PRIMARY KEY (source_id, post_id)
        );`, tableName)

	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AddRecord(post *model.Post, user *model.User) error {
	tableName := getTableName(user)
	query := fmt.Sprintf("INSERT INTO %s (platform, source_id, author_name, post_id) VALUES ($1, $2, $3, $4);", tableName)
	if _, err := r.db.Exec(query, post.PlatformName, post.SourceId, post.Author, post.PostId); err != nil {
		return err
	}
	return nil
}

func getTableName(user *model.User) string {
	return user.Id + user.PlatformName.String()
}
