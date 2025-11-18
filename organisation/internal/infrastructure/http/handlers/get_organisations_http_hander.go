package handlers

import (
	"encoding/json"
	"net/http"
	"organisation/internal/application/queries"
	"organisation/internal/application/queries/handlers"
	"organisation/internal/domain/entities"
)

type GetOrganisationsHttpDto struct {
	Id          string  `json:"id"`
	LegalName   string  `json:"legalName"`
	TradingName string  `json:"tradingName"`
	Website     *string `json:"website"`
}

func NewGetOrganisationsHttpDto(organisation entities.Organisation) GetOrganisationsHttpDto {
	return GetOrganisationsHttpDto{
		Id:          organisation.Id,
		LegalName:   organisation.LegalName,
		TradingName: organisation.TradingName,
		Website:     organisation.Website,
	}
}

type GetOrganisationsHttpHandler struct {
	queryHandler handlers.GetOrganisationsQueryHandler
}

func NewGetOrganisationsHttpHandler(queryHandler handlers.GetOrganisationsQueryHandler) *GetOrganisationsHttpHandler {
	return &GetOrganisationsHttpHandler{
		queryHandler: queryHandler,
	}
}

func (handler *GetOrganisationsHttpHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	organisations, _ := handler.queryHandler.Handle(queries.GetOrganisationsQuery{
		Page: new(int),
		Size: new(int),
	})

	organisationsDto := make([]GetOrganisationsHttpDto, 0, len(organisations))

	for _, organisation := range organisations {
		organisationsDto = append(organisationsDto, NewGetOrganisationsHttpDto(organisation))
	}

	organisationsJson, _ := json.Marshal(organisationsDto)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(organisationsJson))
}
