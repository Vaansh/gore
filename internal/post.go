package internal

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
