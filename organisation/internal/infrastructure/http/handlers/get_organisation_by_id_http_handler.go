package handlers

import (
	"encoding/json"
	"net/http"
	"organisation/internal/application/queries"
	"organisation/internal/application/queries/handlers"
	"organisation/internal/domain/entities"

	"github.com/gorilla/mux"
)

type GetOrganisationByIdHttpDto struct {
	Id          string  `json:"id"`
	LegalName   string  `json:"legalName"`
	TradingName string  `json:"tradingName"`
	Website     *string `json:"website"`
}

func NewGetOrganisationByIdHttpDto(organisation entities.Organisation) GetOrganisationByIdHttpDto {
	return GetOrganisationByIdHttpDto{
		Id:          organisation.Id,
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
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
