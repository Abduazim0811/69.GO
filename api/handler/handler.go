package handler

import (
	"Tasks/internal/models"
	"Tasks/internal/postgres"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *postgres.Task
}



func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{
		taskService: postgres.NewTask(db),
	}
}
// @Router 			/tasks [post]
// @Summary			CREATE TASKS
// @Description 	This method creates tasks
// @Security		BearerAuth
// @Tags			TASK
// @Accept 			json
// @Produce			json
// @Param 			body body models.Task true "Task"
// @Success			201  string gin.H
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			500  {object} models.StandartError
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println(task.Title)
	createdTask, err := h.taskService.StoreCreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": createdTask.ID})
}

// @Router 			/tasks [get]
// @Summary			GET TASKS
// @Description 	This method gets tasks
// @Security		BearerAuth
// @Tags			TASK Tag
// @Accept 			json
// @Produce			json
// @Param 			id query int true "ID"
// @Success			201  {object} []models.Task
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			500  {object} models.StandartError
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.taskService.StoreGetbyIdTasks(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// @Router 			/tasks [get]
// @Summary			GET TASKS
// @Description 	This method gets tasks
// @Security		BearerAuth
// @Tags			TASK Tag
// @Produce			json
// @Success			201  {object} []models.Task
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			500  {object} models.StandartError
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskService.StoreGetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
// @Router 			/tasks [put]
// @Summary			put TASKS
// @Description 	This method put tasks
// @Security		BearerAuth
// @Tags			TASK Tag
// @Accept 			json
// @Produce			json
// @Param 			id query int true "ID"
// @Success			201  {object} models.Task
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			500  {object} models.StandartError
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	task.ID = int32(id)

	updatedTask, err := h.taskService.StoreUpdateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}
// @Router 			/tasks [delete]
// @Summary			delte TASKS
// @Description 	This method delete tasks
// @Security		BearerAuth
// @Tags			TASK Tag
// @Accept 			json
// @Produce			json
// @Param 			id query int true "ID"
// @Success			201  {object} map[string]any
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			500  {object} models.StandartError
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := h.taskService.StoreDeleteTask(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
