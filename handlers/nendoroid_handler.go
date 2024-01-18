package handlers

import (
	r "figures/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NendoroidHandler struct {
    repo *r.NendoroidRepository
}

func Init() *NendoroidHandler {
    repo := r.Init()
    
    return &NendoroidHandler{
        repo: repo,
    }
}

func (h *NendoroidHandler) GetHomePage(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "Hi there")
}

func (h *NendoroidHandler) GetAllNendoroids(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, h.repo.GetAllNendoroids())
}

func (h *NendoroidHandler) GetNendoroidById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Fatal(err)
    }
    nendo, err := h.repo.GetNendoroidById(id)
    if err != nil {
        c.JSON(404, gin.H{"message": "Nendoroid not found"})
    } else {
        c.IndentedJSON(http.StatusOK, nendo)
    }
}
