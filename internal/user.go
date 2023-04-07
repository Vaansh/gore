package internal

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
