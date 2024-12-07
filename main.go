package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type todo struct {
	ID        uuid.UUID `json:"id"`
	Item      string    `json:"item"`
	Completed bool      `json:"completed"`
}

var todos = []todo{
	{
		ID:        uuid.New(),
		Item:      "Clean Room",
		Completed: false,
	},
	{
		ID:        uuid.New(),
		Item:      "Cooking",
		Completed: false,
	},
	{
		ID:        uuid.New(),
		Item:      "Watch movies",
		Completed: true,
	},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:12345")
}
