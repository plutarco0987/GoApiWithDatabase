package main

import (
	"GoApiWithDataBase/internal/database"
	routes "GoApiWithDataBase/internal/routes"
	"log"
)

func main() {
	database.Connect()
	r := routes.SetupRoutes(database.DB)

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
