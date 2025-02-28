package services

import (
	"net/http"
	"net/url"
	"os"
	"robotica_concursos/models"

	"github.com/gin-gonic/gin"
)

// RegistroService defines the interface for the registration service
type RegistroService interface {
	RegisterParticipante(c *gin.Context)
	VerifyParticipante(c *gin.Context)
}

// registroService is the implementation of RegistroService
type registroService struct{}

// NewRegistroService creates a new instance of RegistroService
func NewRegistroService() RegistroService {
	return &registroService{}
}

func (s *registroService) RegisterParticipante(c *gin.Context) {
	var participante models.Participante
	if err := c.ShouldBindJSON(&participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := GenerateRandomToken()
	participante.TokenVerificacion = token
	verificationURL := "http://" + os.Getenv("IP_PUBLICA_SERVER") + ":" + os.Getenv("PORT_SERVER") + "/verify?email=" + url.QueryEscape(participante.Correo) + "&token=" + url.QueryEscape(token)
	NewSMTPEmailSender().SendEmail(
		participante.Correo,
		"Correo de verificacion",
		"Para verificar correo haga click en el siguiente enlace: "+verificationURL)

	db := GetDatabase()
	if err := db.Create(&participante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	participante.Password = ""
	c.JSON(http.StatusOK, gin.H{"data": participante})
}

func (s *registroService) VerifyParticipante(c *gin.Context) {
	email := c.Query("email")
	token := c.Query("token")

	if email == "" || token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and token are required"})
		return
	}

	db := GetDatabase()
	var participante models.Participante
	if err := db.Where("correo = ? AND token_verificacion = ?", email, token).First(&participante).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participante not found"})
		return
	}

	if participante.Verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario ya a sido verificado"})
		return
	}

	participante.Verified = true
	if err := db.Save(&participante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
