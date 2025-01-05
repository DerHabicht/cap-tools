package pgrepo

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/derhabicht/cap-tools/pkg/units"
)

func (repo *PGRepository) CreateOrUpdateUnit(unit units.Unit) error {
	q := `
INSERT INTO units (charter_number, kind, category, name, address, city)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (charter_number) DO UPDATE
SET kind = $2, category = $3, name = $4, address = $5, city = $6;
`

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = stmt.Exec(
		unit.CharterNumber(),
		unit.Kind(),
		unit.Category(),
		unit.Name(),
		unit.Address(),
		unit.City(),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *PGRepository) listAllUnits(offset, limit int) ([]units.Unit, error) {
	q := `
SELECT charter_number, kind, category, name, address, city
FROM units
ORDER BY charter_number
OFFSET $1 LIMIT $2;
`

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := stmt.Query(offset, limit)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var res []units.Unit
	for rows.Next() {
		var charterNumber units.UnitCharterNumber
		var kind units.UnitKind
		var category units.UnitCategory
		var name string
		var address string
		var city string

		err = rows.Scan(&charterNumber, &kind, &category, &name, &address, &city)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		u := units.NewUnit(
			charterNumber,
			kind,
			category,
			name,
			address,
			city,
		)

		res = append(res, u)
	}

	return res, nil
}

func (repo *PGRepository) searchUnits(searchName string, offset, limit int) ([]units.Unit, error) {
	q := `
SELECT charter_number, kind, category, name, address, city
FROM units
WHERE LOWER(name) LIKE $1
ORDER BY charter_number
OFFSET $2 LIMIT $3;
`

	search := fmt.Sprintf("%%%s%%", strings.ToLower(searchName))

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := stmt.Query(search, offset, limit)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var res []units.Unit
	for rows.Next() {
		var charterNumber units.UnitCharterNumber
		var kind units.UnitKind
		var category units.UnitCategory
		var name string
		var address string
		var city string

		err = rows.Scan(&charterNumber, &kind, &category, &name, &address, &city)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		u := units.NewUnit(
			charterNumber,
			kind,
			category,
			name,
			address,
			city,
		)

		res = append(res, u)
	}

	return res, nil
}

func (repo *PGRepository) ListUnits(searchName string, offset, limit int) ([]units.Unit, error) {
	if searchName == "" {
		return repo.listAllUnits(offset, limit)
	}

	return repo.searchUnits(searchName, offset, limit)
}

func (repo *PGRepository) FetchUnit(charterNumber units.UnitCharterNumber) (units.Unit, error) {
	q := `
SELECT charter_number, kind, category, name, address, city
FROM units
WHERE charter_number = $1;
`
	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return units.Unit{}, errors.WithStack(err)
	}

	res, err := stmt.Query(charterNumber)
	if err != nil {
		return units.Unit{}, errors.WithStack(err)
	}

	var unit units.Unit

	found := res.Next()
	if !found {
		return units.Unit{}, errors.Errorf("unit %s not found", charterNumber)
	}

	var cn units.UnitCharterNumber
	var kind units.UnitKind
	var category units.UnitCategory
	var name string
	var address string
	var city string

	err = res.Scan(&cn, &kind, &category, &name, &address, &city)
	if err != nil {
		return units.Unit{}, errors.WithStack(err)
	}

	unit = units.NewUnit(
		cn,
		kind,
		category,
		name,
		address,
		city,
	)

	return unit, nil
}

func (repo *PGRepository) DeleteUnit(unit units.Unit) error {
	q := `DELETE FROM units WHERE charter_number = $1;`

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = stmt.Exec(unit.CharterNumber())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
