package capa5

import (
	"time"
)

type MeetingBlock struct {
	start    time.Time
	end      time.Time
	topic    string
	location string
}

func NewMeetingBlock(
	start time.Time,
	end time.Time,
	topic string,
	location string,
) *MeetingBlock {
	return &MeetingBlock{
		start:    start,
		end:      end,
		topic:    topic,
		location: location,
	}
}
