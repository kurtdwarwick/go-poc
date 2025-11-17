package handlers

import (
	"encoding/json"
	"net/http"
	"organisation/internal/application/queries/handlers"
	"organisation/internal/domain/entities"
	"organisation/internal/domain/queries"

	"github.com/gorilla/mux"
)

type GetOrganisationByIdHttpDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewGetOrganisationByIdHttpDto(organisation entities.Organisation) GetOrganisationByIdHttpDto {
	return GetOrganisationByIdHttpDto{
		Id:   organisation.Id,
		Name: organisation.Name,
	}
}

type GetOrganisationByIdHttpHandler struct {
	queryHandler handlers.GetOrganisationByIdHandler
}

func NewGetOrganisationByIdHttpHandler(queryHandler handlers.GetOrganisationByIdHandler) *GetOrganisationByIdHttpHandler {
	return &GetOrganisationByIdHttpHandler{
		queryHandler: queryHandler,
	}
}

func (handler *GetOrganisationByIdHttpHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)

	organisationId := variables["id"]

	organisation, error := handler.queryHandler.Handle(queries.GetOrganisationByIdQuery{
		Id: organisationId,
	})

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if organisation == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	organisationJson, _ := json.Marshal(NewGetOrganisationByIdHttpDto(*organisation))

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(organisationJson))
}
