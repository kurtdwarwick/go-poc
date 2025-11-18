package handlers

import (
	"organisation/internal/application/queries"
	"organisation/internal/application/repositories"
	"organisation/internal/domain/entities"
)

type GetOrganisationByIdHandler struct {
	repository repositories.OrganisationRepository
}

func NewGetOrganisationByIdHandler(repository repositories.OrganisationRepository) *GetOrganisationByIdHandler {
	return &GetOrganisationByIdHandler{
		repository: repository,
	}
}

func (handler *GetOrganisationByIdHandler) Handle(query queries.GetOrganisationByIdQuery) (*entities.Organisation, error) {
	return handler.repository.GetOrganisationById(query.Id)
}
