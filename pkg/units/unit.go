package units

type Unit struct {
	charterNumber UnitCharterNumber
	kind          UnitKind
	category      UnitCategory
	name          string
}

func NewUnit(
	charterNumber UnitCharterNumber,
	kind UnitKind,
	category UnitCategory,
	name string,
) Unit {
	return Unit{
		charterNumber: charterNumber,
		kind:          kind,
		category:      category,
		name:          name,
	}
}

func (u Unit) CharterNumber() UnitCharterNumber {
	return u.charterNumber
}

func (u Unit) Kind() UnitKind {
	return u.kind
}

func (u Unit) Category() UnitCategory {
	return u.category
}

func (u Unit) Name() string {
	return u.name
}
