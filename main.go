package main

import (
	"net/http"
    r "figures/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    h := Init()

    router.GET("/", h.getHomePage)
    router.GET("/nendoroid", h.getAllNendoroids)
    router.GET("/nendoroid/:id", h.getNendoroidById)

    router.Run("localhost:8080")
}

type NendoroidHandler struct {
    repo *r.NendoroidRepository
}

func Init() *NendoroidHandler {
    repo := r.Init()
    
    return &NendoroidHandler{
        repo: repo,
    }
}

func (h *NendoroidHandler) getHomePage(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "Hi there")
}

func (h *NendoroidHandler) getAllNendoroids(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, h.repo.GetAllNendoroids())
}

func (h *NendoroidHandler) getNendoroidById(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "cool") //h.repo.Get)
}
