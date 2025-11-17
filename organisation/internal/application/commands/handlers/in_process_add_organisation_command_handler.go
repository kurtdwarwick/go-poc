package handlers

import (
	"organisation/internal/application/repositories"
	"organisation/internal/domain/commands"
	"organisation/internal/domain/entities"
)

type AddOrganisationCommandHandler struct {
	repository repositories.OrganisationRepository
}

func NewAddOrganisationCommandHandler(repository repositories.OrganisationRepository) *AddOrganisationCommandHandler {
	return &AddOrganisationCommandHandler{
		repository: repository,
	}
}

func (handler *AddOrganisationCommandHandler) Handle(command commands.CreateOrganisationCommand) (*string, error) {
	organisationId, error := handler.repository.AddOrganisation(entities.Organisation{
		Name: command.Name,
	})

	return organisationId, error
}
