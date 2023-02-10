package interfaces

import "github.com/testresumeguru/models"

type PersonalInfo interface {
	create(models.PersonalInfo) bool
}