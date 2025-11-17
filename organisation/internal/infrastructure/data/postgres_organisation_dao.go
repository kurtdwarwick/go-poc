package data

import (
	"context"
	"errors"
	"log"
	"organisation/internal/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOrganisationDAO struct {
	db *gorm.DB
}

func NewPostgresOrganisationDAO() *PostgresOrganisationDAO {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=organisation sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entities.Organisation{})

	return &PostgresOrganisationDAO{
		db: db,
	}
}

func (dao *PostgresOrganisationDAO) GetOrganisations() ([]entities.Organisation, error) {
	context := context.Background()
	organisations, error := gorm.G[entities.Organisation](dao.db).Find(context)

	if error != nil {
		return nil, error
	}

	return organisations, nil
}

func (dao *PostgresOrganisationDAO) GetOrganisationById(id string) (*entities.Organisation, error) {
	context := context.Background()
	organisation, error := gorm.G[entities.Organisation](dao.db).Where("id = ?", id).Limit(1).Find(context)

	if error != nil {
		return nil, error
	}

	if len(organisation) == 0 {
		return nil, errors.New("organisation not found")
	}

	return &organisation[0], nil
}

func (dao *PostgresOrganisationDAO) CreateOrganisation(
	organisation *entities.Organisation,
	callback func(organisationId string, organisation *entities.Organisation) error) error {
	context := context.Background()

	error := dao.db.Transaction(func(tx *gorm.DB) error {
		error := gorm.G[entities.Organisation](dao.db).Create(context, organisation)

		if error != nil {
			return error
		}

		error = callback(organisation.Id, organisation)

		if error != nil {
			return error
		}

		return nil
	})

	if error != nil {
		return error
	}

	return nil
}
