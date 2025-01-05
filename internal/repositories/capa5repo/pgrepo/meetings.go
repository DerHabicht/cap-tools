package pgrepo

import (
	"github.com/google/uuid"

	"github.com/derhabicht/cap-tools/pkg/capa5"
	"github.com/derhabicht/cap-tools/pkg/units"
)

func (repo *PGRepository) CreateOrUpdateMeeting(meeting capa5.Meeting) error {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) ListMeetings(charterNumber units.UnitCharterNumber, offset, limit int) ([]capa5.Meeting, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) FetchMeeting(meetingID uuid.UUID) (capa5.Meeting, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) DeleteMeeting(meeting capa5.Meeting) error {
	//TODO implement me
	panic("implement me")
}
