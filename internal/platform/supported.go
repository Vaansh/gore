package platform

type Name string

const (
	YOUTUBE   Name = "YT"
	INSTAGRAM Name = "IG"
)

func (n Name) String() string {
	if n == YOUTUBE {
		return "YT"
	} else if n == INSTAGRAM {
		return "IG"
	} else {
		return ""
	}
}
