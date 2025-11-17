package handlers

import (
	"encoding/json"
	"net/http"
	"organisation/internal/application/queries/handlers"
	"organisation/internal/domain/entities"
	"organisation/internal/domain/queries"
)

type GetOrganisationsHttpDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewGetOrganisationsHttpDto(organisation entities.Organisation) GetOrganisationsHttpDto {
	return GetOrganisationsHttpDto{
		Id:   organisation.Id,
		Name: organisation.Name,
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
