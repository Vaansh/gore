package main

type Post struct {
	ID, sourceLink       string
	tag, caption, author string
	platformName         main.PlatformName
}

func NewPost(id, caption, author string) *Post {
	return &Post{
		ID:           id,
		caption:      caption,
		author:       author,
		platformName: PLATFORM,
	}
}
