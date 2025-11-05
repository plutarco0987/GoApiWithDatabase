package routes

import (
	"GoApiWithDataBase/internal/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userHandler := &handlers.UserHandler{DB: db}

	api := router.Group("/api")
	{
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return router
}
