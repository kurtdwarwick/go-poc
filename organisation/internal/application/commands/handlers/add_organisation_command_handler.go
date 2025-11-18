package handlers

import (
	"log"
	"organisation/internal/application/repositories"
	"organisation/internal/domain/commands"
	"organisation/internal/domain/entities"
	"organisation/internal/domain/events"
	"shared"
)

type AddOrganisationCommandHandler struct {
	eventPublisher shared.EventPublisher
	repository     repositories.OrganisationRepository
}

func NewAddOrganisationCommandHandler(eventPublisher shared.EventPublisher, repository repositories.OrganisationRepository) *AddOrganisationCommandHandler {
	return &AddOrganisationCommandHandler{
		eventPublisher: eventPublisher,
		repository:     repository,
	}
}

func (handler *AddOrganisationCommandHandler) Handle(command commands.CreateOrganisationCommand) (*string, error) {
	organisationId, error := handler.repository.AddOrganisation(
		entities.Organisation{
			LegalName:   command.LegalName,
			TradingName: command.TradingName,
			Website:     &command.Website,
		},
		func(organisationId string, organisation *entities.Organisation) error {
			event := &events.OrganisationAddedEvent{
				Id:          organisationId,
				LegalName:   command.LegalName,
				TradingName: command.TradingName,
				Website:     command.Website,
			}

			error := handler.eventPublisher.Publish(event)

			return error
		})

	if error != nil {
		log.Printf("Failed to add organisation: %v", error)
		return nil, error
	}

	if error != nil {
		log.Printf("Failed to publish event: %v", error)
		return nil, error
	}

	return organisationId, error
}
