package main

import (
	"net/http"
    r "figures/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    repo := r.Init()
    router.GET("/", getHomePage)
    router.GET("/nendoroid", getAllNendoroids)
    router.GET("/nendoroid/:id", getNendoroidById)

    router.Run("localhost:8080")
}

func getHomePage(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "Hi there")
}

func getAllNendoroids(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, repo.GetAllNendoroids())
}

func getNendoroidById(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "cool id bro")
}
