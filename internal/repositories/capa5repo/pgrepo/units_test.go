package pgrepo

import (
	"github.com/stretchr/testify/assert"

	"github.com/derhabicht/cap-tools/pkg/units"
)

func (s *PGRepoTestSuite) TestUnits() {
	cn, err := units.ParseCharterNumber("RMR-UT-080")
	assert.NoError(s.T(), err)

	unit := units.NewUnit(
		cn,
		units.Squadron,
		units.CadetUnit,
		"Blackhawk Cadet Squadron",
		"12953 S Minuteman Dr",
		"Draper, UT 84020",
	)

	err = s.repo.CreateOrUpdateUnit(unit)
	assert.NoError(s.T(), err)

	res, err := s.repo.FetchUnit(cn)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), unit.CharterNumber(), res.CharterNumber())
	assert.Equal(s.T(), unit.Kind(), res.Kind())
	assert.Equal(s.T(), unit.Category(), res.Category())
	assert.Equal(s.T(), unit.Name(), res.Name())
	assert.Equal(s.T(), unit.Address(), res.Address())
	assert.Equal(s.T(), unit.City(), res.City())

	unit = units.NewUnit(
		unit.CharterNumber(),
		units.Group,
		units.CompositeUnit,
		"Holy Blackhawk Cadet Empire",
		unit.Address(),
		unit.City(),
	)

	err = s.repo.CreateOrUpdateUnit(unit)
	assert.NoError(s.T(), err)

	res, err = s.repo.FetchUnit(cn)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), unit.CharterNumber(), res.CharterNumber())
	assert.Equal(s.T(), unit.Kind(), res.Kind())
	assert.Equal(s.T(), unit.Category(), res.Category())
	assert.Equal(s.T(), unit.Name(), res.Name())
	assert.Equal(s.T(), unit.Address(), res.Address())
	assert.Equal(s.T(), unit.City(), res.City())

	list, err := s.repo.ListUnits("", 0, 100)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), len(list), 1)

	err = s.repo.DeleteUnit(unit)
	assert.NoError(s.T(), err)

	list, err = s.repo.ListUnits("", 0, 100)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), len(list), 0)

	_, err = s.repo.FetchUnit(cn)
	assert.Error(s.T(), err)
}

func (s *PGRepoTestSuite) TestUnitsSearch() {
	cn07, err := units.ParseCharterNumber("RMR-UT-007")
	assert.NoError(s.T(), err)
	sl := units.NewUnit(
		cn07,
		units.Squadron,
		units.SeniorUnit,
		"Salt Lake Senior Squadron",
		"640 N 2360 W",
		"Salt Lake City, UT 84116",
	)
	err = s.repo.CreateOrUpdateUnit(sl)
	assert.NoError(s.T(), err)

	cn80, err := units.ParseCharterNumber("RMR-UT-080")
	assert.NoError(s.T(), err)
	blackhawk := units.NewUnit(
		cn80,
		units.Squadron,
		units.CadetUnit,
		"Blackhawk Cadet Squadron",
		"12953 S Minuteman Dr",
		"Draper, UT 84020",
	)
	err = s.repo.CreateOrUpdateUnit(blackhawk)
	assert.NoError(s.T(), err)

	res, err := s.repo.ListUnits("blackhawk", 0, 10)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), len(res), 1)
	assert.Equal(s.T(), res[0].CharterNumber(), blackhawk.CharterNumber())
	assert.Equal(s.T(), res[0].Kind(), blackhawk.Kind())
	assert.Equal(s.T(), res[0].Category(), blackhawk.Category())
	assert.Equal(s.T(), res[0].Name(), blackhawk.Name())
	assert.Equal(s.T(), res[0].Address(), blackhawk.Address())
	assert.Equal(s.T(), res[0].City(), blackhawk.City())
}
