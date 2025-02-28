package services

import (
	"robotica_concursos/models"

	"gorm.io/gorm"
)

// ParticipanteServiceInterface defines the methods that any implementation of ParticipanteService should have
type ParticipanteServiceInterface interface {
	GetParticipantes() ([]models.Participante, error)
	GetParticipante(id string) (models.Participante, error)
	CreateParticipante(participante models.Participante) (models.Participante, error)
	UpdateParticipante(participante models.Participante) (models.Participante, error)
	DeleteParticipante(id string) error
}

type ParticipanteService struct {
	db *gorm.DB
}

func NewParticipanteService(db *gorm.DB) *ParticipanteService {
	return &ParticipanteService{db: db}
}

func (ps ParticipanteService) GetParticipantes() ([]models.Participante, error) {
	var participantes []models.Participante
	if err := ps.db.Find(&participantes).Error; err != nil {
		return nil, err
	}
	return participantes, nil
}

func (ps ParticipanteService) GetParticipante(id string) (models.Participante, error) {
	var participante models.Participante
	if err := ps.db.First(&participante, id).Error; err != nil {
		return models.Participante{}, err
	}
	return participante, nil
}

func (ps ParticipanteService) CreateParticipante(participante models.Participante) (models.Participante, error) {
	if err := ps.db.Create(&participante).Error; err != nil {
		return models.Participante{}, err
	}
	return participante, nil
}

func (ps ParticipanteService) UpdateParticipante(participante models.Participante) (models.Participante, error) {
	if err := ps.db.Save(&participante).Error; err != nil {
		return models.Participante{}, err
	}
	return participante, nil
}

func (ps ParticipanteService) DeleteParticipante(id string) error {
	if err := ps.db.Delete(&models.Participante{}, id).Error; err != nil {
		return err
	}
	return nil
}
