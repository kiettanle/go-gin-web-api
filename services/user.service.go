package services

import "go-gin-web-api/models"

type UserService interface {
	Create(*models.User) error
	Get(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(*models.User) error
	Delete(*string) error
}
