package repositories

import (
	"organisation/internal/domain/data"
	"organisation/internal/domain/entities"
	"organisation/internal/domain/policies"

	"shared"

	"github.com/google/uuid"
)

type OrganisationRepository struct {
	organisationDao data.OrganisationDAO
	organisations   map[string]entities.Organisation

	policyHandler *shared.PolicyHandler
}

func NewOrganisationRepository(organisationDao data.OrganisationDAO) *OrganisationRepository {

	policyHandler := shared.NewPolicyHandler(policies.OrganisationNamePolicy{})

	return &OrganisationRepository{
		organisationDao: organisationDao,
		organisations:   make(map[string]entities.Organisation),
		policyHandler:   policyHandler,
	}
}

func (repository *OrganisationRepository) GetOrganisations() ([]entities.Organisation, error) {
	organisations, error := repository.organisationDao.GetOrganisations()

	return organisations, error
}

func (repository *OrganisationRepository) GetOrganisationById(id string) (*entities.Organisation, error) {
	organisation, error := repository.organisationDao.GetOrganisationById(id)

	return organisation, error
}

func (repository *OrganisationRepository) AddOrganisation(
	organisation entities.Organisation,
	callback func(organisationId string, organisation *entities.Organisation) error) (*string, error) {
	organisationId := uuid.New().String()

	error := repository.policyHandler.Validate(organisation)

	if error != nil {
		return nil, error
	}

	organisation.Id = organisationId

	error = repository.organisationDao.CreateOrganisation(&organisation, callback)

	return &organisationId, error
}
