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
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	for _, a := range equipos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipo not found"})
}

func PostEquipo(c *gin.Context) {
	services.PostEquipo(c)
}

func UpdateEquipo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var updatedEquipo Equipo

	if err := c.BindJSON(&updatedEquipo); err != nil {
		return
	}

	for i, a := range equipos {
		if a.ID == id {
			equipos[i] = updatedEquipo
			c.IndentedJSON(http.StatusOK, updatedEquipo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipo not found"})
}

func DeleteEquipo(c *gin.Context) {
	DeleteEquipo(c)
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/equipos", GetEquipos)
	router.GET("/equipos/:id", GetEquipoByID)
	router.POST("/equipos", PostEquipo)
	router.PUT("/equipos/:id", UpdateEquipo)
	router.DELETE("/equipos/:id", DeleteEquipo)
}
