package cap

type UnitCategory int

const (
	UnknownUnitCategory UnitCategory = iota
	AdminUnit
	CompositeUnit
	CadetUnit
	SeniorUnit
)
