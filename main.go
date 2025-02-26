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
	router.POST("/users", addUser)
	router.GET("/users/:id", getUserId)
	router.Run("localhost:8080")

}
func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
func addUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
func getUserId(c *gin.Context) {
	id := c.Param("id")
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}
