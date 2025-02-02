package controllers

import (
	"net/http"
	"strconv"

	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type Equipo struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

var equipos = []Equipo{
	{ID: 1, Name: "Equipo1"},
	{ID: 2, Name: "Equipo2"},
}

func GetEquipos(c *gin.Context) {
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
	services.GetEquipos(c, page, tam)
}

func GetEquipoByID(c *gin.Context) {
	services.GetEquipoByID(c)
}

func PostEquipo(c *gin.Context) {
	services.PostEquipo(c)
}

func UpdateEquipo(c *gin.Context) {
	services.UpdateEquipo(c)
}

func DeleteEquipo(c *gin.Context) {
	services.DeleteEquipo(c)
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/equipos", GetEquipos)
	router.GET("/equipos/:id", GetEquipoByID)
	router.POST("/equipos", PostEquipo)
	router.PUT("/equipos/:id", UpdateEquipo)
	router.DELETE("/equipos/:id", DeleteEquipo)
}
