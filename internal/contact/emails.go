package contact

import (
	"fmt"

	"github.com/ag7if/go-files"
	set "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"

	"github.com/derhabicht/cap-tools/internal/clients"
	"github.com/derhabicht/cap-tools/internal/config"
	"github.com/derhabicht/cap-tools/pkg"
)

func excluded(cn cap.UnitCharterNumber) bool {
	excludedUnits := config.ExcludedUnits()

	return excludedUnits.Contains(cn)
}

func buildEmailList(members []pkg.Member) string {
	emailSet := set.NewSet[string]()

	for _, member := range members {
		if !excluded(member.Unit.CharterNumber()) {
			continue
		}

		for _, contact := range member.Contacts {
			if contact.ContactPriority == pkg.Primary &&
				(contact.ContactType == pkg.Email || contact.ContactType == pkg.CadetParentEmail) {

				emailSet.Add(contact.Contact)
			}
		}
	}

	list := ""
	for _, contact := range emailSet.ToSlice() {
		list = fmt.Sprintf("%s%s\n", list, contact)
	}

	return list
}

func ScrapeEmails(report files.File) error {
	members, err := clients.ParseMemberContactReport(report)
	if err != nil {
		return errors.WithStack(err)
	}

	emailList := buildEmailList(members)
	fmt.Println(emailList)

	return nil
}
