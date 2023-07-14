package api

import (
	"github.com/Vaansh/gore/internal/platform"
	"time"
)

type RunInstagramTaskRequest struct {
	IgUserId             string          `json:"igUserId"`
	LongLivedAccessToken string          `json:"lAccessToken"`
	Hashtags             string          `json:"igPostTags"`
	PublisherIds         []string        `json:"publisherIds"`
	Sources              []platform.Name `json:"sources"`
	SubscriberId         string          `json:"subscriberId"`
	Frequency            time.Duration   `json:"frequency,string"`
}
