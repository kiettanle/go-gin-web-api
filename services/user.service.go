//go:generate mockgen --source user.service.go --destination mock/user.service_mock.go  --package mock

package services

import "go-gin-web-api/models"

type UserService interface {
	Create(*models.User) (*models.User, error)
	Get(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(*string, *models.User) (*models.User, error)
	Delete(*string) error
}
