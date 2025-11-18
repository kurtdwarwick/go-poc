package handlers

import (
	"organisation/internal/application/queries"
	"organisation/internal/application/repositories"
	"organisation/internal/domain/entities"
)

type GetOrganisationsQueryHandler struct {
	repository repositories.OrganisationRepository
}

func NewGetOrganisationsQueryHandler(repository repositories.OrganisationRepository) *GetOrganisationsQueryHandler {
	return &GetOrganisationsQueryHandler{
		repository: repository,
	}
}

func (handler *GetOrganisationsQueryHandler) Handle(query queries.GetOrganisationsQuery) ([]entities.Organisation, error) {
	return handler.repository.GetOrganisations()
}
