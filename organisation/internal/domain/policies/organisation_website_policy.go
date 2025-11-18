package policies

import (
	"errors"
	"organisation/internal/domain/entities"
	"regexp"
)

type OrganisationWebsitePolicy struct{}

func (policy OrganisationWebsitePolicy) Validate(value any) error {
	organisation, ok := value.(entities.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.Website != nil {
		regex, _ := regexp.Compile(`^(https?://)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`)

		if !regex.MatchString(*organisation.Website) {
			return errors.New("organisation website is not a valid URL")
		}
	}

	return nil
}
