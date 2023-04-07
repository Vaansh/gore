package main

type PlatformName string

const (
	INSTAGRAM PlatformName = "IG"
)

type DestinationPost interface{}

type InstagramPost struct {
	ID, caption, author string
	sourceLink          string
	platformName        PlatformName
}

func NewInstagramPost(id, caption, author string) InstagramPost {
	return InstagramPost{
		ID:           id,
		caption:      caption,
		author:       author,
		platformName: INSTAGRAM,
	}
}
