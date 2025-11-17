package policies

import (
	"organisation/internal/domain/entities"
	"testing"
)

func TestOrganisationNamePolicy_Validate(t *testing.T) {
	policy := OrganisationNamePolicy{}

	t.Run("should return an error if the organisation name is empty", func(t *testing.T) {
		organisation := entities.Organisation{
			Name: "",
		}

		err := policy.Validate(organisation)

		if err == nil {
			t.Errorf("expected an error")
		}
	})

	t.Run("should not return an error if the organisation name is not empty", func(t *testing.T) {
		organisation := entities.Organisation{
			Name: "Test Organisation",
		}

		err := policy.Validate(organisation)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
