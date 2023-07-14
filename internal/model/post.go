package model

import (
	"fmt"
	"github.com/Vaansh/gore/internal/platform"
)

type Post struct {
	PostId, SourceId     string
	tag, Caption, Author string
	PlatformName         platform.Name
}

func NewPost(id, caption, author, sourceId string, platformName platform.Name) *Post {
	return &Post{
		PostId:       id,
		Caption:      caption,
		Author:       author,
		SourceId:     sourceId,
		PlatformName: platformName,
	}
}

func (p Post) GetParams() (string, string, platform.Name, string) {
	return p.PostId, p.Author, p.PlatformName, p.Caption
}

func (p Post) String() string {
	return fmt.Sprintf("Id: %s Author:%s Platform:%s", p.PostId, p.Author, p.PlatformName)
}
