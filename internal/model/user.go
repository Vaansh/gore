package model

import "github.com/Vaansh/gore/internal/platform"

type User struct {
	ID           string
	platformName platform.Name
}

func NewUser(ID string, platform platform.Name) *User {
	return &User{
		ID:           ID,
		platformName: platform,
	}
}
