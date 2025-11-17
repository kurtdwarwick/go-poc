package events

type OrganisationAddedEvent struct {
	Id   string
	Name string
}

func (event OrganisationAddedEvent) GetType() string {
	return "organisation.added"
}
