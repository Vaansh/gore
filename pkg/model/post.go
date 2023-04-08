package model

import (
	"github.com/Vaansh/gore/pkg"
)

type User struct {
	ID           string
	platformName pkg.PlatformName
}

func NewUser(ID string, platform pkg.PlatformName) *User {
	return &User{
		ID:           ID,
		platformName: platform,
	}
}
