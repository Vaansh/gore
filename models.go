package main

type PlatformName string

const (
	PLATFORM PlatformName = "PF"
)

type User struct {
	ID           string
	platformName PlatformName
}

func NewUser(id, platform string) *User {
	return &User{
		ID:           id,
		platformName: PLATFORM,
	}
}

type Post struct {
	ID, sourceLink       string
	tag, caption, author string
	platformName         PlatformName
}

func NewPost(id, caption, author string) *Post {
	return &Post{
		ID:           id,
		caption:      caption,
		author:       author,
		platformName: PLATFORM,
	}
}
