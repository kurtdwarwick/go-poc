package policies

import (
	"organisation/internal/domain/entities"
	"testing"
)

// Unit tests should be focused on predictable behaviour (business logic) and not on implementation details, and should be focused
// on quality of tests, rather than quantity.

func TestOrganisationLegalNamePolicy_Validate(t *testing.T) {
	policy := OrganisationLegalNamePolicy{}

	t.Run("should return an error if the organisation legal name is empty", func(t *testing.T) {
		organisation := entities.Organisation{
			LegalName: "",
		}

		err := policy.Validate(organisation)

		if err == nil {
			t.Errorf("expected an error")
		}
	})

	t.Run("should not return an error if the organisation legal name is not empty", func(t *testing.T) {
		organisation := entities.Organisation{
			LegalName: "Test Organisation",
		}

		err := policy.Validate(organisation)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
