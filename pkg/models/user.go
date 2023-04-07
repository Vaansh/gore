package models

import (
	"github.com/Vaansh/gore/pkg"
)

type Post struct {
	ID, sourceLink       string
	tag, caption, author string
	platformName         pkg.PlatformName
}

func NewPost(id, caption, author string) *Post {
	return &Post{
		ID:           id,
		caption:      caption,
		author:       author,
		platformName: pkg.PLATFORM,
	}
}
