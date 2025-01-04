package gormrepo

import (
	"gorm.io/gorm/clause"

	"github.com/derhabicht/cap-tools/pkg/units"
)

type unit struct {
	model
	CharterNumber units.UnitCharterNumber `gorm:"unique"`
	Kind          *units.UnitKind
	Category      *units.UnitCategory
	Name          string
	Address       string
	City          string
}

func newUnit(u units.Unit) unit {
	var kind *units.UnitKind
	if u.Kind() != units.UnknownUnitKind {
		*kind = u.Kind()
	}

	var category *units.UnitCategory
	if u.Category() != units.UnknownUnitCategory {
		*category = u.Category()
	}

	return unit{
		CharterNumber: u.CharterNumber(),
		Kind:          kind,
		Category:      category,
		Name:          u.Name(),
		Address:       u.Address(),
		City:          u.City(),
	}
}

func (u unit) ToUnit() units.Unit {
	var kind units.UnitKind
	if u.Kind == nil {
		kind = units.UnknownUnitKind
	} else {
		kind = *u.Kind
	}

	var category units.UnitCategory
	if u.Category == nil {
		category = units.UnknownUnitCategory
	} else {
		category = *u.Category
	}

	return units.NewUnit(
		u.CharterNumber,
		kind,
		category,
		u.Name,
		u.Address,
		u.City,
	)
}

func (repo *GormRepository) CreateOrUpdateUnit(u units.Unit) (units.Unit, error) {
	ru := newUnit(u)

	/*
		db.Clauses(clause.OnConflict{
		  Columns:   []clause.Column{{Name: "id"}},
		  DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("GREATEST(count, VALUES(count))")}),
		}).Create(&users)
	*/

	upsert := clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
	}

}
