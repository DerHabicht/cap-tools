package units

type Unit struct {
	charterNumber UnitCharterNumber
	kind          UnitKind
	category      UnitCategory
	name          string
	address       string
	city          string
}

func NewUnit(
	charterNumber UnitCharterNumber,
	kind UnitKind,
	category UnitCategory,
	name string,
	address string,
	city string,
) Unit {
	return Unit{
		charterNumber: charterNumber,
		kind:          kind,
		category:      category,
		name:          name,
		address:       address,
		city:          city,
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

func (u Unit) Address() string {
	return u.address
}

func (u Unit) City() string {
	return u.city
}
