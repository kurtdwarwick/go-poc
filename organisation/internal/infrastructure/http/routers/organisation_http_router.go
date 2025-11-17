package routers

import (
	"organisation/internal/infrastructure/http/handlers"

	"github.com/gorilla/mux"
)

type OrganisationHttpRouter struct {
	getOrganisationByIdHandler handlers.GetOrganisationByIdHttpHandler
	getOrganisationsHandler    handlers.GetOrganisationsHttpHandler

	addOrganisationHandler handlers.AddOrganisationHttpHandler
}

func NewOrganisationHttpRouter(getOrganisationByIdHandler handlers.GetOrganisationByIdHttpHandler, getOrganisationsHandler handlers.GetOrganisationsHttpHandler, addOrganisationHandler handlers.AddOrganisationHttpHandler) *OrganisationHttpRouter {
	return &OrganisationHttpRouter{
		getOrganisationByIdHandler: getOrganisationByIdHandler,
		getOrganisationsHandler:    getOrganisationsHandler,
		addOrganisationHandler:     addOrganisationHandler,
	}
}

func (router *OrganisationHttpRouter) RegisterRoutes(r *mux.Router) {
	organisationsRoute := r.PathPrefix("/organisations").Subrouter()

	organisationsRoute.HandleFunc("", router.getOrganisationsHandler.Handle).Methods("GET")
	organisationsRoute.HandleFunc("/{id}", router.getOrganisationByIdHandler.Handle).Methods("GET")
	organisationsRoute.HandleFunc("", router.addOrganisationHandler.Handle).Methods("POST")
}
