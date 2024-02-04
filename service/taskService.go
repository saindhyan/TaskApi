package service

import (
	"database/sql"
	"net/http"
	"strconv"
	Entity "task/Entity"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect(data *sql.DB) {
	db = data
}

func CreateTask(c *gin.Context) {
	var task Entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, erre := db.Exec("INSERT INTO task (title ,description ,dueDate,status) VALUES (?,?,?,?)", task.Title, task.Description, task.DueDate, task.Status)
	if erre != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erre.Error()})
		return
	}

	id, _ := result.LastInsertId()
	task.ID = int(id)

	c.JSON(http.StatusCreated, task)

}

func GetTask(c *gin.Context) {
	var tasks []Entity.Task

	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	for rows.Next() {
		var task Entity.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(200, tasks)
}

func GetTaskbyid(c *gin.Context) {
	id := c.Param("id")
	var task Entity.Task
	err := db.QueryRow("SELECT * FROM task WHERE id=?", id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, task)
}
func UpdateTask(c *gin.Context) {

	id := c.Param("id")
	var task Entity.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE task SET title=? ,description=?,dueDate=?,status=? WHERE id=?", task.Title, task.Description, task.DueDate, task.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	taskId, err := strconv.Atoi(id)
	if err == nil {
		task.ID = taskId
	}
	c.JSON(http.StatusOK, task)
}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM task WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Massage": "Deleted"})
}
