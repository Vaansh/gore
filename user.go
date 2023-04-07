package main

type User struct {
	ID           string
	platformName main.PlatformName
}

func NewUser(id, platform string) *User {
	return &User{
		ID:           id,
		platformName: PLATFORM,
	}
}
