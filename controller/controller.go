package controller

import (
	"task/service"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.POST("/tasks", service.CreateTask)
	r.GET("/tasks/:id", service.GetTaskbyid)
	r.GET("/tasks", service.GetTask)
	r.PUT("/tasks/:id", service.UpdateTask)
	r.DELETE("/tasks/:id", service.DeleteTask)

	return r
}
