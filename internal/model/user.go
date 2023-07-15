package model

import "github.com/Vaansh/gore/internal/platform"

type User struct {
	Id           string        `json:"id"`
	PlatformName platform.Name `json:"platformName"`
}

func NewUser(ID string, platform platform.Name) *User {
	return &User{
		Id:           ID,
		PlatformName: platform,
	}
}
