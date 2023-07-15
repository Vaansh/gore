package gore

type Platform string

const (
	YOUTUBE   Platform = "yt"
	INSTAGRAM Platform = "ig"
)

func (n Platform) String() string {
	if n == YOUTUBE {
		return "yt"
	} else if n == INSTAGRAM {
		return "ig"
	} else {
		return ""
	}
}

func (n Platform) OfficialName() string {
	if n == YOUTUBE {
		return "YouTube"
	} else if n == INSTAGRAM {
		return "Instagram"
	} else {
		return ""
	}
}
