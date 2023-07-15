package model

import (
	"fmt"
	"github.com/Vaansh/gore"
)

type Post struct {
	PostId       string         `json:"postId"`
	SourceId     string         `json:"sourceId"`
	Tag          string         `json:"tag"`
	Caption      string         `json:"caption"`
	Author       string         `json:"author"`
	PlatformName go_pubsub.Name `json:"platformName"`
}

func NewPost(id, caption, author, sourceId string, platformName go_pubsub.Name) *Post {
	return &Post{
		PostId:       id,
		Caption:      caption,
		Author:       author,
		SourceId:     sourceId,
		PlatformName: platformName,
	}
}

func (p Post) GetParams() (string, string, go_pubsub.Name, string) {
	return p.PostId, p.Author, p.PlatformName, p.Caption
}

func (p Post) String() string {
	return fmt.Sprintf("Id: %s Author:%s Platform:%s", p.PostId, p.Author, p.PlatformName)
}
