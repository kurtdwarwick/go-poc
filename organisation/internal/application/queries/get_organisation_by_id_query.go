package queries

type GetOrganisationByIdQuery struct {
	Id string
}

func NewGetOrganisationByIdQuery(id string) *GetOrganisationByIdQuery {
	return &GetOrganisationByIdQuery{
		Id: id,
	}
}
