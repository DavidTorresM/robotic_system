package main

import (
	"log"
	"robotica_concursos/controllers"
	"robotica_concursos/models"
	"robotica_concursos/services"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Iniciando sistema de concursos robots")

	log.Println("Generando DDL de la base de datos")
	services.ConnectDatabase()

	err := models.MigrateTables(services.DB)
	if err != nil {
		log.Fatal("Error generando el DDL de la base de datos:", err)
		return
	}
	db := services.GetDatabase()
	models.InsertCategorias(db)

	//prendiendo el servidor de gin

	router := gin.Default()

	controllers.RegisterRoutes(router)
	controllers.RegisterRoutesRobots(router)
	controllers.RegisterParticipanteRoutes(router)
	controllers.RegisterLoginRoutes(router)
	controllers.RegisterRegistreRoutes(router)

	router.Run(os.Getenv("IP_SERVER") + ":" + os.Getenv("PORT_SERVER"))

}
