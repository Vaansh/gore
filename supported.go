package go_pubsub

type Name string

const (
	YOUTUBE   Name = "yt"
	INSTAGRAM Name = "ig"
)

func (n Name) FullName() string {
	if n == YOUTUBE {
		return "YouTube"
	} else if n == INSTAGRAM {
		return "Instagram"
	} else {
		return ""
	}
}

func (n Name) String() string {
	if n == YOUTUBE {
		return "yt"
	} else if n == INSTAGRAM {
		return "ig"
	} else {
		return ""
	}
}
