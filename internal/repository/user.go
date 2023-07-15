package repository

import "github.com/Vaansh/gore/internal/model"

// Repository pattern

type UserRepository interface {
	AddRecord(tableName string, post *model.Post) error
	CheckIfRecordExists(tableName string, post *model.Post) (bool, error)
}
