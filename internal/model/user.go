package model

import (
	"github.com/Vaansh/gore"
)

type User struct {
	Id           string        `json:"id"`
	PlatformName gore.Platform `json:"platformName"`
}

func NewUser(ID string, platform gore.Platform) *User {
	return &User{
		Id:           ID,
		PlatformName: platform,
	}
}
