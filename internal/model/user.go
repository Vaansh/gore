package model

import (
	"github.com/Vaansh/gore"
)

type User struct {
	Id           string         `json:"id"`
	PlatformName go_pubsub.Name `json:"platformName"`
}

func NewUser(ID string, platform go_pubsub.Name) *User {
	return &User{
		Id:           ID,
		PlatformName: platform,
	}
}
