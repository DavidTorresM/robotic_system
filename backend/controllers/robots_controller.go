package controllers

import (
	"net/http"
	"strconv"

	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type RobotsController struct {
	service services.RobotService
}

func NewRobotsController(service services.RobotService) *RobotsController {
	return &RobotsController{service: service}
}

func (rc *RobotsController) GetRobots(c *gin.Context) {
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
	rc.service.GetRobots(c, page, tam)
}

func (rc *RobotsController) GetRobotByID(c *gin.Context) {
	rc.service.GetRobotByID(c)
}

func (rc *RobotsController) PostRobot(c *gin.Context) {
	rc.service.PostRobot(c)
}

func (rc *RobotsController) UpdateRobot(c *gin.Context) {
	rc.service.UpdateRobot(c)
}

func (rc *RobotsController) DeleteRobot(c *gin.Context) {
	rc.service.DeleteRobot(c)
}

func RegisterRoutesRobots(router *gin.Engine, service services.RobotService) {
	controller := NewRobotsController(service)
	router.GET("/robots", controller.GetRobots)
	router.GET("/robots/:id", controller.GetRobotByID)
	router.POST("/robots", controller.PostRobot)
	router.PUT("/robots/:id", controller.UpdateRobot)
	router.DELETE("/robots/:id", controller.DeleteRobot)
}
