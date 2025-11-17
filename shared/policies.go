package shared

type Policy interface {
	Validate(value any) error
}

type PolicyHandler struct {
	Policies []Policy
}

func NewPolicyHandler(policies ...Policy) *PolicyHandler {
	return &PolicyHandler{
		Policies: policies,
	}
}

func (handler *PolicyHandler) Validate(value any) error {
	for _, policy := range handler.Policies {
		if err := policy.Validate(value); err != nil {
			return err
		}
	}

	return nil
}
