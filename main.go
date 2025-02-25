package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
	DNI  int32  `json:"dni"`
}

var users = []user{
	{ID: "1", Name: "Dante", Mail: "dantetest@gmail.com", DNI: 123456789},
	{ID: "2", Name: "Sofia", Mail: "sofietest@gmail,com", DNI: 789456123},
	{ID: "3", Name: "Jhon", Mail: "jhon@gmail.com", DNI: 456789123},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUser)
	router.Run("localhost:8080")
}
func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user)
}
