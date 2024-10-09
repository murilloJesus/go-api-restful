package controllers

import (
	"fmt"
	"net/http"
	"restful/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
	{ID: 1, Title: "Fazer arroz", Completed: false},
	{ID: 2, Title: "Limpar o jardim", Completed: false},
}

var lastId = 2

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastId++
	newTask.ID = lastId
	tasks = append(tasks, newTask)
	c.JSON(http.StatusOK, newTask)
}

func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not Found"})
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Task %s deleted", task.Title)})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
