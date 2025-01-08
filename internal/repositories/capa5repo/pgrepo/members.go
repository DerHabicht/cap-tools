package pgrepo

import (
	"github.com/pkg/errors"

	"github.com/derhabicht/cap-tools/pkg/members"
	"github.com/derhabicht/cap-tools/pkg/units"
)

func (repo *PGRepository) CreateOrUpdateMember(member members.Member) error {
	q := `
INSERT INTO members (capid, member_category, last_name, first_name, grade, unit_charter)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (capid) DO UPDATE
SET member_category = $2, last_name = $3, first_name = $4, grade = $5, unit_charter = $6;
`
	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = stmt.Exec(
		member.CAPID(),
		member.MemberCategory(),
		member.LastName(),
		member.FirstName(),
		member.Grade(),
		member.Unit().CharterNumber(),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *PGRepository) ListMembers(charterNumber units.UnitCharterNumber, offset, limit int) ([]members.Member, error) {
	unit, err := repo.FetchUnit(charterNumber)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := `
SELECT capid, member_category, last_name, first_name, grade
FROM members
WHERE unit_charter = $1
ORDER BY last_name
OFFSET $2 LIMIT $3;
`

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := stmt.Query(charterNumber, offset, limit)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var res []members.Member
	for rows.Next() {
		var capidRes uint
		var memberCategory members.MemberCategory
		var lastName string
		var firstName string
		var grade members.Grade

		err = rows.Scan(&capidRes, &memberCategory, &lastName, &firstName, &grade)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		member := members.NewMember(
			capidRes,
			memberCategory,
			lastName,
			firstName,
			grade,
			unit,
		)

		res = append(res, member)
	}

	return res, nil
}

func (repo *PGRepository) FetchMember(capid uint) (members.Member, error) {
	q := `
SELECT capid, member_category, last_name, first_name, grade, unit_charter
FROM members
WHERE capid = $1;
`

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return members.Member{}, errors.WithStack(err)
	}

	res, err := stmt.Query(capid)
	if err != nil {
		return members.Member{}, errors.WithStack(err)
	}

	found := res.Next()
	if !found {
		err = ErrNotFound{Object: "members", Key: capid}
		return members.Member{}, errors.WithStack(err)
	}

	var capidRes uint
	var memberCategory members.MemberCategory
	var lastName string
	var firstName string
	var grade members.Grade
	var charterNumber units.UnitCharterNumber

	err = res.Scan(&capidRes, &memberCategory, &lastName, &firstName, &grade, &charterNumber)
	if err != nil {
		return members.Member{}, errors.WithStack(err)
	}

	unit, err := repo.FetchUnit(charterNumber)
	if err != nil {
		return members.Member{}, errors.WithStack(err)
	}

	member := members.NewMember(
		capidRes,
		memberCategory,
		lastName,
		firstName,
		grade,
		unit,
	)

	return member, nil
}

func (repo *PGRepository) DeleteMember(member members.Member) error {
	q := `
DELETE FROM members
WHERE capid = $1;
`

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return errors.WithStack(err)
	}

	res, err := stmt.Exec(member.CAPID())
	if err != nil {
		return errors.WithStack(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errors.WithStack(err)
	}

	if affected == 0 {
		return errors.WithStack(ErrNotFound{Object: "members", Key: member.CAPID()})
	}

	return nil
}
