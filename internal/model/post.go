package model

import (
	"fmt"
	"github.com/Vaansh/gore/internal/platform"
)

type Post struct {
	ID, sourceLink       string
	tag, Caption, Author string
	PlatformName         platform.Name
}

func NewPost(id, caption, author string, platformName platform.Name) *Post {
	return &Post{
		ID:           id,
		Caption:      caption,
		Author:       author,
		PlatformName: platformName,
	}
}

func (p Post) GetParams() (string, string, platform.Name, string) {
	return p.ID, p.Author, p.PlatformName, p.Caption
}

func (p Post) String() string {
	return fmt.Sprintf("ID: %s Author:%s Platform:%s", p.ID, p.Author, p.PlatformName)
}
