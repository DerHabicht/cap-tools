package clients

import (
	"encoding/csv"
	"strconv"

	"github.com/ag7if/cap/v2"
	"github.com/ag7if/go-files"
	"github.com/pkg/errors"

	"github.com/derhabicht/cap-tools/pkg"
)

const (
	orgCol             = 0
	capidCol           = 1
	rankCol            = 2
	fullNameCol        = 3
	mbrTypeCol         = 4
	contactTypeCol     = 5
	contactPriorityCol = 6
	contactInfoCol     = 7
	contactNameCol     = 8
)

func openReport(f files.File) ([][]string, error) {
	osf, err := f.Open()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer osf.Close()

	r := csv.NewReader(osf)

	records, err := r.ReadAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return records, nil
}

func parseContactRow(record []string) (*pkg.Member, pkg.MemberContact, error) {
	org, err := cap.ParseCharterNumber(record[orgCol])
	if err != nil {
		return nil, pkg.MemberContact{}, errors.WithStack(err)
	}

	capidint, err := strconv.Atoi(record[capidCol])
	if err != nil {
		return nil, pkg.MemberContact{}, errors.WithStack(err)
	}
	capid := uint(capidint)

	grade, err := cap.ParseGrade(record[rankCol])
	if err != nil {
		return nil, pkg.MemberContact{}, errors.WithStack(err)
	}

	first, middle, last := pkg.ParseMemberName(record[fullNameCol])

	mbrType, err := cap.ParseMemberType(record[mbrTypeCol])
	if err != nil {
		return nil, pkg.MemberContact{}, errors.WithStack(err)
	}

	contactType, err := pkg.ParseContactType(record[contactTypeCol])
	if err != nil {
		return nil, pkg.MemberContact{}, errors.WithStack(err)
	}

	contactPriority, err := pkg.ParseContactPriority(record[contactPriorityCol])
	if err != nil {
		return nil, pkg.MemberContact{}, errors.WithStack(err)
	}

	contactInfo := record[contactInfoCol]

	contactName := record[contactNameCol]

	memberContact := pkg.MemberContact{
		Contact:         contactInfo,
		ContactName:     contactName,
		ContactPriority: contactPriority,
		ContactType:     contactType,
	}

	member := &pkg.Member{
		CAPID:      capid,
		Grade:      grade,
		MemberType: mbrType,
		Unit:       cap.NewUnit(org, cap.UnknownUnitKind, cap.UnknownUnitCategory, ""),
		LastName:   last,
		MiddleName: middle,
		FirstName:  first,
		Contacts:   []pkg.MemberContact{memberContact},
	}

	return member, memberContact, nil
}

func parseMembers(records [][]string) (map[uint]*pkg.Member, error) {
	members := make(map[uint]*pkg.Member)

	for _, record := range records {
		mbr, contact, err := parseContactRow(record)
		if err != nil {
			// TODO: Log warn message in this case
			continue
		}

		_, ok := members[mbr.CAPID]
		if ok {
			members[mbr.CAPID].Contacts = append(members[mbr.CAPID].Contacts, contact)
		} else {
			members[mbr.CAPID] = mbr
		}
	}

	return members, nil
}

func ParseMemberContactReport(f files.File) ([]pkg.Member, error) {
	records, err := openReport(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	memberMap, err := parseMembers(records)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var members []pkg.Member

	for _, member := range memberMap {
		members = append(members, *member)
	}

	return members, nil
}
