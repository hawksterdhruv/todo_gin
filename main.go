package main

import (
	"log"
	"net/http"

	"todo_gin/models"

	"github.com/gin-gonic/gin"
)

// var todos = []Todo{
// 	{Id: "1", Title: "Walmart", Description: "Shop for groceries"},
// 	{Id: "2", Title: "LeetCode", Description: "Solve problems on LeetCode"},
// }

var db = models.InitDb()

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}

func getTodos(c *gin.Context) {
	var todos []models.Todo
	db.Find(&todos)
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo models.Todo

	err := c.BindJSON(&newTodo)
	if err != nil {
		log.Print("Couldn't convert the incoming data into a Todo", err)
	}

	result := db.Create(&newTodo)
	if result.Error != nil {
		log.Print("== Error inserting record ==", result.Error)
		c.IndentedJSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	// todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/", getHello)
	router.GET("/todo/list", getTodos)
	router.POST("/todo", addTodo)

	models.InitDb()

	router.Run("localhost:8000")
}
