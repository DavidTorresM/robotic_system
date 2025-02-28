package controllers

import (
	"net/http"
	"robotica_concursos/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService *services.LoginService
}

func NewLoginController(loginService *services.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
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

	participante, err := ctrl.loginService.Authenticate(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
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

func RegisterLoginRoutes(router *gin.Engine, loginService *services.LoginService) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	lc := NewLoginController(loginService)
	router.POST("/login", lc.Login)
}
