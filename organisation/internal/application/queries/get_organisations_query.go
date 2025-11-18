package queries

type GetOrganisationsQuery struct {
	Page *int
	Size *int
}

func NewGetOrganisationsQuery(page *int, size *int) *GetOrganisationsQuery {
	return &GetOrganisationsQuery{
		Page: page,
		Size: size,
	}
}
