package handlers

import (
	"encoding/json"
	"net/http"
	"organisation/internal/application/commands/handlers"
	"organisation/internal/domain/commands"
)

type AddOrganisationHttpDto struct {
	Name string `json:"name"`
}

type AddOrganisationHttpHandler struct {
	commandHandler handlers.AddOrganisationCommandHandler
}

func NewAddOrganisationHttpHandler(commandHandler handlers.AddOrganisationCommandHandler) *AddOrganisationHttpHandler {
	return &AddOrganisationHttpHandler{
		commandHandler: commandHandler,
	}
}

func (handler *AddOrganisationHttpHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	var organisation AddOrganisationHttpDto

	error := json.NewDecoder(request.Body).Decode(&organisation)

	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	event, error := handler.commandHandler.Handle(commands.CreateOrganisationCommand{
		Name: organisation.Name,
	})

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(writer).Encode(map[string]string{
			"error": error.Error(),
		})

		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(map[string]string{
		"id": event.Id,
	})
}
