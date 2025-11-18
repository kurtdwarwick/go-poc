package policies

import (
	"errors"
	"organisation/internal/domain/entities"
)

type OrganisationLegalNamePolicy struct{}

func (policy OrganisationLegalNamePolicy) Validate(value any) error {
	organisation, ok := value.(entities.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.LegalName == "" {
		return errors.New("organisation legal name is required")
	}

	return nil
}
