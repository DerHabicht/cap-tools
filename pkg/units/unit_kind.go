package units

type UnitKind int

const (
	UnknownUnitKind UnitKind = iota
	Group
	Squadron
	Flight
)
