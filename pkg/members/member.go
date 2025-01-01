package members

type Member struct {
	capid        uint
	lastName     string
	firstName    string
	grade        Grade
	officeSymbol string
}

func NewMember(
	capid uint,
	lastName string,
	firstName string,
	grade Grade,
	officeSymbol string,
) Member {
	return Member{
		capid:        capid,
		lastName:     lastName,
		firstName:    firstName,
		grade:        grade,
		officeSymbol: officeSymbol,
	}
}
