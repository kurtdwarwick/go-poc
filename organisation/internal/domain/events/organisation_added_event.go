package events

type OrganisationAddedEvent struct {
	Id          string
	LegalName   string
	TradingName string
	Website     string
}

func (event OrganisationAddedEvent) GetType() string {
	return "organisation.added"
}
