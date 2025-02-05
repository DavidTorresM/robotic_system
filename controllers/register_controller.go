package controllers

import (
	"net/http"
	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type Participante struct {
	ID       uint    `gorm:"primaryKey"`
	Nombre   string  `gorm:"size:100;not null"`
	Correo   string  `gorm:"size:100;unique;not null"`
	Password string  `gorm:"size:100;" json:"Password,omitempty"`
	Telefono string  `gorm:"size:15"`
	EquipoID uint    `gorm:"not null"`
	Equipo   *Equipo `gorm:"foreignKey:EquipoID" json:"Equipo,omitempty"`
}

type Equipo struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
}

func RegisterParticipante(c *gin.Context) {
	var participante Participante
	if err := c.ShouldBindJSON(&participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := services.GetDatabase()
	if err := db.Create(&participante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": participante})
}

func RegisterRegistreRoutes(router *gin.Engine) {
	router.POST("/register", RegisterParticipante)
}
