package repository

import (
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore"
	"github.com/Vaansh/gore/internal/model"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB, id string, platform gore.Platform) (*PostgresUserRepository, error) {
	r := &PostgresUserRepository{db: db}
	if err := r.createUser(model.NewUser(id, platform)); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *PostgresUserRepository) createUser(user *model.User) error {
	tableName := getTableName(user)
	query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
			post_id VARCHAR(20) NOT NULL,
			source_id VARCHAR(50) NOT NULL,
			author_name VARCHAR(100),
            platform VARCHAR(5) NOT NULL,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            PRIMARY KEY (source_id, post_id)
        );`, tableName)
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) AddRecord(tableName string, post *model.Post) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (platform, source_id, author_name, post_id) 
		VALUES ($1, $2, $3, $4);`, tableName)
	if _, err := r.db.Exec(query, post.PlatformName, post.SourceId, post.Author, post.PostId); err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) CheckIfRecordExists(tableName string, post *model.Post) (bool, error) {
	var count int
	query := fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM %s 
		WHERE platform = $1 AND source_id = $2 AND post_id = $3;`, tableName)
	if err := r.db.QueryRow(query, post.PlatformName, post.SourceId, post.PostId).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func getTableName(user *model.User) string {
	return user.PlatformName.String() + "_" + user.Id
}
