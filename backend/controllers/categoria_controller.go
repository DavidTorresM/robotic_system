package controllers

import (
	"net/http"
	"strconv"

	"robotica_concursos/models"
	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type CategoriaController struct {
	service services.CategoriaService
}

func NewCategoriaController(service services.CategoriaService) *CategoriaController {
	return &CategoriaController{service: service}
}

func (cc *CategoriaController) GetCategorias(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	categorias, err := cc.service.GetCategorias(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categorias)
}

func (cc *CategoriaController) GetCategoria(c *gin.Context) {
	id := c.Param("id")
	categoria, err := cc.service.GetCategoriaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria not found"})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func (cc *CategoriaController) CreateCategoria(c *gin.Context) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.service.CreateCategoria(categoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, categoria)
}

func (cc *CategoriaController) UpdateCategoria(c *gin.Context) {
	id := c.Param("id")
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.service.UpdateCategoria(id, categoria); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria not found"})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func (cc *CategoriaController) DeleteCategoria(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.service.DeleteCategoria(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func RegisterCategoriaRoutes(router *gin.Engine, service services.CategoriaService) {
	cc := NewCategoriaController(service)
	router.GET("/categorias", cc.GetCategorias)
	router.GET("/categorias/:id", cc.GetCategoria)
	router.POST("/categorias", cc.CreateCategoria)
	router.PUT("/categorias/:id", cc.UpdateCategoria)
	router.DELETE("/categorias/:id", cc.DeleteCategoria)
}
