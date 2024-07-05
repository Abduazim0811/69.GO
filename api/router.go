package api

import (
	"Tasks/api/handler"
	"database/sql"
	"fmt"
	"log"

	// _ "Tasks/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title	Project: Swagger Intro
// @description This swagger UI was created in lesson
// @version 1.0
// @securityDefinitions.apikey Bearerauth
// @in header
// @name Authorization
func Api(db *sql.DB, port string) {

	taskhandler := handler.NewTaskHandler(db)

	router := gin.Default()

	router.POST("/tasks", taskhandler.CreateTask)
	router.GET("/tasks/:id", taskhandler.GetTaskByID)
	// router.GET("/tasks", taskhandler.GetTasks)
	// router.PUT("/tasks/:id", taskhandler.UpdateTask)
	// router.DELETE("tasks/:id", taskhandler.DeleteTask)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Printf("Starting server on port %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
