package main

import (
	"log"
	"robotica_concursos/controllers"
	"robotica_concursos/models"
	"robotica_concursos/services"
	"time"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Iniciando sistema de concursos robots")
	log.Println("Cargando variables de entorno")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
		os.Exit(0)
	}

	log.Println("Generando DDL de la base de datos")
	services.ConnectDatabase()

	err = models.MigrateTables(services.GetDatabase())
	if err != nil {
		log.Fatal("Error generando el DDL de la base de datos:", err)
		return
	}
	db := services.GetDatabase()
	models.InsertCategorias(db)

	//prendiendo el servidor de gin

	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permitir cualquier origen (cambia esto en producción)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	controllers.RegisterRoutes(router)
	controllers.RegisterRoutesRobots(router)
	controllers.RegisterParticipanteRoutes(router)
	controllers.RegisterLoginRoutes(router)
	controllers.RegisterRegistreRoutes(router)
	controllers.RegisterCategoriaRoutes(router)
	controllers.RegisterRoutesCompeticion(router)

	router.Run(os.Getenv("IP_SERVER") + ":" + os.Getenv("PORT_SERVER"))

}
