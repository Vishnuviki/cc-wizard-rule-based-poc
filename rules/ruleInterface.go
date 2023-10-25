package services

import (
	"connectivity-checker-wizard/models"
)

type Rule interface {
	Execute(request any) models.ResponseData
	SetNextRule(nextRule Rule)
}
