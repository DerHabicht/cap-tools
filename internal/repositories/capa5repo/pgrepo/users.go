package pgrepo

import (
	"github.com/google/uuid"

	"github.com/derhabicht/cap-tools/pkg/capa5"
	"github.com/derhabicht/cap-tools/pkg/units"
)

func (repo *PGRepository) CreateOrUpdateUser(user capa5.User) error {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) ListUsers(charterNumber units.UnitCharterNumber, offset, limit int) ([]capa5.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) FetchUser(id uuid.UUID) (capa5.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) DeleteUser(user capa5.User) error {
	//TODO implement me
	panic("implement me")
}
