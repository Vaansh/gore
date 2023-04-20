package model

import (
	"github.com/Vaansh/gore/internal"
)

type User struct {
	ID           string
	platformName internal.PlatformName
}

func NewUser(ID string, platform internal.PlatformName) *User {
	return &User{
		ID:           ID,
		platformName: platform,
	}
}
