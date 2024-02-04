package controller

import (
	"task/service"

	"github.com/gin-gonic/gin"
)

func Routes()*gin.Engine {
	r := gin.Default()
	r.POST("/task", service.CreateTask)
	r.GET("/task/:id", service.GetTaskbyid)
	r.GET("/task", service.GetTask)
	r.PUT("/task/:id", service.UpdateTask)
	r.DELETE("/task/:id", service.DeleteTask)

	return r
}
