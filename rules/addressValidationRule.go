package services

import (
	"connectivity-checker-wizard/models"
)

type AddressValidationRule struct {
	Name     string
	NextRule Rule // may be we can make it a List
}

func (r *AddressValidationRule) Execute(request any) models.ResponseData {
	// TODO - Execute current rule

	// communicate with API server if needed

	// based on the situation we can chain the next rule
	if r.NextRule != nil {
		r.NextRule.Execute(nil)
	}
	return models.ResponseData{}
}

func (r *AddressValidationRule) SetNextRule(nextRule Rule) {
	// this case it's - nil
}
