package data

import (
	"errors"
	"organisation/internal/domain/entities"
)

type InMemoryOrganisationDAO struct {
	organisations map[string]entities.Organisation
}

func NewInMemoryOrganisationDAO() *InMemoryOrganisationDAO {
	return &InMemoryOrganisationDAO{
		organisations: make(map[string]entities.Organisation),
	}
}

func (dao *InMemoryOrganisationDAO) GetOrganisations() ([]entities.Organisation, error) {
	organisations := make([]entities.Organisation, 0, len(dao.organisations))

	for _, organisation := range dao.organisations {
		organisations = append(organisations, organisation)
	}

	return organisations, nil
}

func (dao *InMemoryOrganisationDAO) GetOrganisationById(id string) (*entities.Organisation, error) {
	organisation, ok := dao.organisations[id]

	if !ok {
		return nil, errors.New("organisation not found")
	}

	return &organisation, nil
}

func (dao *InMemoryOrganisationDAO) CreateOrganisation(
	organisation *entities.Organisation,
	callback func(organisationId string, organisation *entities.Organisation) error) error {
	dao.organisations[organisation.Id] = *organisation

	error := callback(organisation.Id, organisation)

	if error != nil {
		return error
	}

	return nil
}
