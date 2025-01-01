package capa5

import (
	"time"

	"github.com/derhabicht/cap-tools/pkg/members"
)

type Meeting struct {
	location         string
	address          string
	start            time.Time
	end              time.Time
	meetingType      MT
	topic            string
	blocks           []MeetingBlock
	uod              []members.Uniform
	additionalOrders string
}

func NewMeeting(
	location string,
	address string,
	start time.Time,
	end time.Time,
	meetingType MT,
	topic string,
	blocks []MeetingBlock,
	uod []members.Uniform,
	additionalOrders string,
) Meeting {
	return Meeting{
		location:         location,
		address:          address,
		start:            start,
		end:              end,
		meetingType:      meetingType,
		topic:            topic,
		blocks:           blocks,
		uod:              uod,
		additionalOrders: additionalOrders,
	}
}
