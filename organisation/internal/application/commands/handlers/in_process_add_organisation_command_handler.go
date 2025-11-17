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

func (handler *AddOrganisationCommandHandler) Handle(command commands.CreateOrganisationCommand) (*events.OrganisationAddedEvent, error) {
	organisationId, error := handler.repository.AddOrganisation(entities.Organisation{
		Name: command.Name,
	})

	if error != nil {
		log.Printf("Failed to add organisation: %v", error)
		return nil, error
	}

	event := &events.OrganisationAddedEvent{
		Id:   *organisationId,
		Name: command.Name,
	}

	error = handler.eventPublisher.Publish(event)

	if error != nil {
		log.Printf("Failed to publish event: %v", error)
		return nil, error
	}

	return event, error
}
