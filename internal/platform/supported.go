package platform

type Name string

const (
	YOUTUBE   Name = "yt"
	INSTAGRAM Name = "ig"
)

func (n Name) String() string {
	if n == YOUTUBE {
		return "yt"
	} else if n == INSTAGRAM {
		return "ig"
	} else {
		return ""
	}
}
