package model

import (
	"github.com/Vaansh/gore/internal/platform"
)

type Post struct {
	ID, sourceLink       string
	tag, caption, author string
	platformName         platform.Name
}

func NewPost(id, caption, author string) *Post {
	return &Post{
		ID:           id,
		caption:      caption,
		author:       author,
		platformName: platform.YOUTUBE,
	}
}
