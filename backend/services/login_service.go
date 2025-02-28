package services

import (
	"errors"
	"robotica_concursos/models"

	"gorm.io/gorm"
)

type LoginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{db: db}
}

func (service *LoginService) Authenticate(username, password string) (*models.Participante, error) {
	var participante models.Participante
	if err := service.db.Where("Correo = ?", username).First(&participante).Error; err != nil {
		return nil, err
	}

	if participante.Password != password {
		return nil, errors.New("invalid username or password")
	}

	return &participante, nil
}
