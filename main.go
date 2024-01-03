package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Tag struct {
}
type Comment struct {
}

type User struct {
}

var todos = []Todo{
	{Id: "1", Title: "Walmart", Description: "Shop for groceries"},
	{Id: "2", Title: "LeetCode", Description: "Solve problems on LeetCode"},
}

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo Todo

	err := c.BindJSON(&newTodo)
	if err != nil {
		log.Print("Couldn't convert the incoming data into a Todo", err)
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/", getHello)
	router.GET("/todo/list", getTodos)
	router.POST("/todo", addTodo)

	router.Run("localhost:8000")
}
