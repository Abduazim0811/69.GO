package api

import (
	"Tasks/api/handler"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Api(db *sql.DB, port string) {
	taskhandler := handler.NewTaskHandler(db)

	router := gin.Default()

	router.POST("/tasks", taskhandler.CreateTask)
	router.GET("/tasks/:id", taskhandler.GetTaskByID)
	router.GET("/tasks", taskhandler.GetTasks)
	router.PUT("/tasks/:id", taskhandler.UpdateTask)
	router.DELETE("tasks/:id", taskhandler.DeleteTask)
	log.Printf("Starting server on port %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
