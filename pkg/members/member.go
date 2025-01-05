package members

type Member struct {
	capid     uint
	lastName  string
	firstName string
	grade     Grade
}

func NewMember(
	capid uint,
	lastName string,
	firstName string,
	grade Grade,
) Member {
	return Member{
		capid:     capid,
		lastName:  lastName,
		firstName: firstName,
		grade:     grade,
	}
}

func (m Member) CAPID() uint {
	return m.capid
}

func (m Member) LastName() string {
	return m.lastName
}

func (m Member) FirstName() string {
	return m.firstName
}

func (m Member) Grade() Grade {
	return m.grade
}
