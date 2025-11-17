package policies

import (
	"errors"
	"organisation/internal/domain/entities"
)

type OrganisationNamePolicy struct{}

func (policy OrganisationNamePolicy) Validate(value any) error {
	organisation, ok := value.(entities.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.Name == "" {
		return errors.New("organisation name is required")
	}

	return nil
}
