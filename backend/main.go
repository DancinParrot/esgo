package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ok")
}

func main() {
	router := gin.Default()
	router.GET("/hello-world", getHelloWorld)

	router.Run("localhost:8080")
}
