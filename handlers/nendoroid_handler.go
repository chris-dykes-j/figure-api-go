package handlers

import (
	r "figures/repositories"
	"net/http"

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
    c.IndentedJSON(http.StatusOK, h.repo.GetNendoroidById(c.Param("id"))) 
}
