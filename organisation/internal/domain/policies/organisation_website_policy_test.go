package policies

import (
	"organisation/internal/domain/entities"
	"testing"
)

// Unit tests should be focused on predictable behaviour (business logic) and not on implementation details, and should be focused
// on quality of tests, rather than quantity.

func TestOrganisationWebsitePolicy_Validate(t *testing.T) {
	policy := OrganisationWebsitePolicy{}

	t.Run("should return an error if the organisation website is not a valid URL", func(t *testing.T) {
		website := "some not valid URL"

		organisation := entities.Organisation{
			Website: &website,
		}

		err := policy.Validate(organisation)

		if err == nil {
			t.Errorf("expected an error")
		}
	})

	t.Run("should not return an error if the organisation website is a valid URL", func(t *testing.T) {
		website := "https://www.example.com"

		organisation := entities.Organisation{
			Website: &website,
		}

		err := policy.Validate(organisation)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("should not return an error if the organisation website is nil", func(t *testing.T) {
		organisation := entities.Organisation{
			Website: nil,
		}

		err := policy.Validate(organisation)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
