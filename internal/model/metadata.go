package model

import (
	"time"
)

type MetaData struct {
	Frequency              time.Duration
	IgUserId               string
	IgPostTags             string
	IgLongLivedAccessToken string
}

func NewInstagramMetaData(igUserId, igLongLivedAccessToken, igPostTags string, frequency time.Duration) *MetaData {
	return &MetaData{
		Frequency:              frequency,
		IgUserId:               igUserId,
		IgPostTags:             igPostTags,
		IgLongLivedAccessToken: igLongLivedAccessToken,
	}
}
