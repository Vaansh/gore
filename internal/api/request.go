package api

import (
	gore "github.com/Vaansh/gore"
	"time"
)

// Api task request dtos live here

type RunInstagramTaskRequest struct {
	IgUserId             string          `json:"igUserId"`
	LongLivedAccessToken string          `json:"lAccessToken"`
	Hashtags             string          `json:"igPostTags"`
	PublisherIds         []string        `json:"publisherIds"`
	Sources              []gore.Platform `json:"sources"`
	SubscriberId         string          `json:"subscriberId"`
	Frequency            time.Duration   `json:"frequency,string"`
}
