package main

import (
	"context"

	"net/http"

	"organisation/internal/application/repositories"

	"organisation/internal/infrastructure/data"
	"organisation/internal/infrastructure/http/routers"

	commandHandlers "organisation/internal/application/commands/handlers"
	queryHandlers "organisation/internal/application/queries/handlers"
	httpHandlers "organisation/internal/infrastructure/http/handlers"

	"github.com/gorilla/mux"
)

func main() {
	organisationDao := data.NewPostgresOrganisationDAO()
	organisationRepository := repositories.NewOrganisationRepository(organisationDao)

	router := mux.NewRouter().StrictSlash(true)

	addOrganisationCommandHandler := commandHandlers.NewAddOrganisationCommandHandler(*organisationRepository)
	getOrganisationsQueryHandler := queryHandlers.NewGetOrganisationsQueryHandler(*organisationRepository)
	getOrganisationByIdQueryHandler := queryHandlers.NewGetOrganisationByIdHandler(*organisationRepository)

	getOrganisationByIdHandler := httpHandlers.NewGetOrganisationByIdHttpHandler(*getOrganisationByIdQueryHandler)
	getOrganisationsHandler := httpHandlers.NewGetOrganisationsHttpHandler(*getOrganisationsQueryHandler)
	addOrganisationHandler := httpHandlers.NewAddOrganisationHttpHandler(*addOrganisationCommandHandler)

	organisationRouter := routers.NewOrganisationHttpRouter(*getOrganisationByIdHandler, *getOrganisationsHandler, *addOrganisationHandler)
	organisationRouter.RegisterRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

	defer server.Shutdown(context.Background())
}
