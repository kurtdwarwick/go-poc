package main

import (
	"context"

	"net/http"

	"organisation/internal/application/repositories"

	"organisation/internal/infrastructure/data"
	"organisation/internal/infrastructure/events/publishers"
	"organisation/internal/infrastructure/http/routers"

	commandHandlers "organisation/internal/application/commands/handlers"
	queryHandlers "organisation/internal/application/queries/handlers"
	httpHandlers "organisation/internal/infrastructure/http/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// You would instaniate the In Memory DAO here, if you need to swap out the implementation. I.e. you could use a Mongo imoplementation or anything else, really.
	organisationDao := data.NewPostgresOrganisationDAO()
	// organisationDao := data.NewInMemoryOrganisationDAO()
	organisationRepository := repositories.NewOrganisationRepository(organisationDao)

	// You would be able to swap this out for any other event publisher implementation.
	eventPublisher := publishers.NewRabbitMQEventPublisher()

	defer eventPublisher.Dispose()

	addOrganisationCommandHandler := commandHandlers.NewAddOrganisationCommandHandler(eventPublisher, *organisationRepository)
	getOrganisationsQueryHandler := queryHandlers.NewGetOrganisationsQueryHandler(*organisationRepository)
	getOrganisationByIdQueryHandler := queryHandlers.NewGetOrganisationByIdHandler(*organisationRepository)

	getOrganisationByIdHandler := httpHandlers.NewGetOrganisationByIdHttpHandler(*getOrganisationByIdQueryHandler)
	getOrganisationsHandler := httpHandlers.NewGetOrganisationsHttpHandler(*getOrganisationsQueryHandler)
	addOrganisationHandler := httpHandlers.NewAddOrganisationHttpHandler(*addOrganisationCommandHandler)

	router := mux.NewRouter().StrictSlash(true)

	organisationRouter := routers.NewOrganisationHttpRouter(*getOrganisationByIdHandler, *getOrganisationsHandler, *addOrganisationHandler)
	organisationRouter.RegisterRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

	defer server.Shutdown(context.Background())
}
