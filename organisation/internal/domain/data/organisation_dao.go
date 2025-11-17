package data

import "organisation/internal/domain/entities"

type OrganisationDAO interface {
	GetOrganisations() ([]entities.Organisation, error)
	GetOrganisationById(id string) (*entities.Organisation, error)

	CreateOrganisation(organisation *entities.Organisation, callback func(organisationId string, organisation *entities.Organisation) error) error
}
