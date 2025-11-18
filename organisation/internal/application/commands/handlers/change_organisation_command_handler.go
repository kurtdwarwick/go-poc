package handlers

import (
	"organisation/internal/application/repositories"
	"organisation/internal/domain/commands"
	"organisation/internal/domain/entities"
	"shared"
)

type ChangeOrganisationCommandHandler struct {
	repository repositories.OrganisationRepository
}

func NewChangeOrganisationCommandHandler(repository repositories.OrganisationRepository) *ChangeOrganisationCommandHandler {
	return &ChangeOrganisationCommandHandler{
		repository: repository,
	}
}

func (handler *ChangeOrganisationCommandHandler) Handle(command commands.ChangeOrganisationCommand) error {
	return handler.repository.ChangeOrganisation(&entities.Organisation{
		Entity: shared.Entity{
			Id: command.OrganisationId,
		},
		LegalName: command.Name,
	}, func(organisation *entities.Organisation) error {
		return nil
	})
}
