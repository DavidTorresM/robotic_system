package controllers

import (
	"net/http"
	"strconv"

	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type EquipoController struct {
	service services.EquipoService
}

func NewEquipoController(service services.EquipoService) *EquipoController {
	return &EquipoController{service: service}
}

func (ec *EquipoController) GetEquipos(c *gin.Context) {
	tam, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid size parameter"})
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid page parameter"})
		return
	}
	ec.service.GetEquipos(c, page, tam)
}

func (ec *EquipoController) GetEquipoByID(c *gin.Context) {
	ec.service.GetEquipoByID(c)
}

func (ec *EquipoController) PostEquipo(c *gin.Context) {
	ec.service.PostEquipo(c)
}

func (ec *EquipoController) UpdateEquipo(c *gin.Context) {
	ec.service.UpdateEquipo(c)
}

func (ec *EquipoController) DeleteEquipo(c *gin.Context) {
	ec.service.DeleteEquipo(c)
}

func RegisterRoutesEquipo(router *gin.Engine, service services.EquipoService) {
	controller := NewEquipoController(service)
	router.GET("/equipos", controller.GetEquipos)
	router.GET("/equipos/:id", controller.GetEquipoByID)
	router.POST("/equipos", controller.PostEquipo)
	router.PUT("/equipos/:id", controller.UpdateEquipo)
	router.DELETE("/equipos/:id", controller.DeleteEquipo)
}
