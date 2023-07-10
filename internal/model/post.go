package model

import (
	"github.com/Vaansh/gore/internal"
)

type Post struct {
	ID, sourceLink       string
	tag, caption, author string
	platformName         internal.PlatformName
}

func NewPost(id, caption, author string) *Post {
	return &Post{
		ID:           id,
		caption:      caption,
		author:       author,
		platformName: internal.YOUTUBE,
	}
}
