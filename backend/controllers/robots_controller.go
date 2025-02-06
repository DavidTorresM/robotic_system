package controllers

import (
	"net/http"
	"strconv"

	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

func GetRobots(c *gin.Context) {
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
	services.GetRobots(c, page, tam)
}

func GetRobotByID(c *gin.Context) {
	services.GetRobotByID(c)
}

func PostRobot(c *gin.Context) {
	services.PostRobot(c)
}

func UpdateRobot(c *gin.Context) {
	services.UpdateRobot(c)
}

func DeleteRobot(c *gin.Context) {
	services.DeleteRobot(c)
}

func RegisterRoutesRobots(router *gin.Engine) {
	router.GET("/robots", GetRobots)
	router.GET("/robots/:id", GetRobotByID)
	router.POST("/robots", PostRobot)
	router.PUT("/robots/:id", UpdateRobot)
	router.DELETE("/robots/:id", DeleteRobot)
}
