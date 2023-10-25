package services

import (
	"connectivity-checker-wizard/models"
)

type SourceNSCNPRule struct {
	Name     string
	NextRule Rule
}

func (r *SourceNSCNPRule) Execute(request any) models.ResponseData {
	// TODO - Execute current rule

	return models.ResponseData{}
}

func (r *SourceNSCNPRule) SetNextRule(nextRule Rule) {

}
