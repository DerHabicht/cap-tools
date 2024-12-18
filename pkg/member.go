package pkg

import (
	"strings"

	"github.com/ag7if/cap/v2"
)

type Member struct {
	CAPID      uint
	Grade      cap.Grade
	MemberType cap.MemberType
	Unit       cap.Unit
	LastName   string
	MiddleName string
	FirstName  string
	Contacts   []MemberContact
}

func ParseMemberName(name string) (first, middle, last string) {
	s := strings.Split(name, " ")

	if len(s) == 3 {
		return s[0], s[1], s[2]
	}

	first = s[0]
	last = s[len(s)-1]
	middle = strings.Join(s[1:len(s)-1], " ")

	return first, middle, last
}
