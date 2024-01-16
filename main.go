package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/", getHomePage)
    router.GET("/nendoroid", getAllNendoroids)
    router.GET("/nendoroid/:id", getNendoroidById)

    router.Run("localhost:8080")
}

func getHomePage(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "Hi there")
}

func getAllNendoroids(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "cool")
}

func getNendoroidById(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "cool id bro")
}
