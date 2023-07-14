package model

import (
	"time"
)

type MetaData struct {
	IgPostTags string
	Frequency  time.Duration
}

func NewInstagramMetaData(igPostTags string, frequency time.Duration) *MetaData {
	return &MetaData{
		IgPostTags: igPostTags,
		Frequency:  frequency,
	}
}
