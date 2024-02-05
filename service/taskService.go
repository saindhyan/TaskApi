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

// service connection with database
func Connect(data *sql.DB) {
	db = data
}

// creating  new task
func CreateTask(c *gin.Context) {
	var task Entity.Task

	// binding json with task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pendingText := "Pending"

	// inserting on database
	result, erre := db.Exec("INSERT INTO task (title ,description ,dueDate,status) VALUES (?,?,?,?)", task.Title, task.Description, task.DueDate, pendingText)
	if erre != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erre.Error()})
		return
	}

	id, _ := result.LastInsertId()

	task.ID = int(id)
	task.Status = pendingText
	c.JSON(http.StatusCreated, task)

}

// getting all task
func GetTask(c *gin.Context) {
	var tasks []Entity.Task

	// geting all tasks from database
	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// itrating over tasks from database
	for rows.Next() {
		var task Entity.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		//make single object
		tasks = append(tasks, task)
	}

	//showing all tasks
	c.JSON(200, tasks)
}

// geting task by id
func GetTaskbyid(c *gin.Context) {
	id := c.Param("id")
	var task Entity.Task

	// geting task from database and binding with json
	err := db.QueryRow("SELECT * FROM task WHERE id=?", id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// showing task
	c.JSON(http.StatusOK, task)
}

// updating task if exists
func UpdateTask(c *gin.Context) {

	id := c.Param("id")
	var task Entity.Task
	// binding json with task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// updated according to json
	result, err := db.Exec("UPDATE task SET title=? ,description=?,dueDate=?,status=? WHERE id=?", task.Title, task.Description, task.DueDate, task.Status, id)
	rowsAffacted, _ := result.RowsAffected()
	if rowsAffacted == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
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

// delete task if exists
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	// deleting task from database
	result, err := db.Exec("DELETE FROM task WHERE id=?", id)
	rowsAffacted, _ := result.RowsAffected()
	if rowsAffacted == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Massage": "Deleted"})
}
