package model

import "github.com/Vaansh/gore/internal/platform"

type User struct {
	Id           string
	PlatformName platform.Name
}

func NewUser(ID string, platform platform.Name) *User {
	return &User{
		Id:           ID,
		PlatformName: platform,
	}
}
