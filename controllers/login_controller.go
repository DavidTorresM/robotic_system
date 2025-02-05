package controllers

import (
	"net/http"
	"robotica_concursos/models"
	"robotica_concursos/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (ctrl *LoginController) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var db = services.GetDatabase()

	var participante models.Participante
	if err := db.Where("Correo = ?", loginData.Username).First(&participante).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if participante.Password != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}

	session := sessions.Default(c)
	session.Set("session_id_part", participante.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	// If authentication is successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func RegisterLoginRoutes(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	lc := NewLoginController()
	router.POST("/login", lc.Login)
}
