package handlers

import (
	r "figures/repositories"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NendoroidHandler struct {
	repo           *r.NendoroidRepository
	supportedLangs [3]string
}

func Init() *NendoroidHandler {
	repo := r.Init()
	supportedLangs := [3]string{"en", "ja", "zh"} // Unsure best place for this.

	return &NendoroidHandler{
		repo:           repo,
		supportedLangs: supportedLangs,
	}
}

func (h *NendoroidHandler) GetHomePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hi there")
}

// TODO Search by Item Name
func (h *NendoroidHandler) GetAllNendoroids(c *gin.Context) {
	lang := c.Query("language")
	if !h.isSupportedLang(lang) {
		lang = "en"
	}
	character := c.Query("character")
	if character == "" {
		character = c.Query("name")
	}

	c.IndentedJSON(http.StatusOK, h.repo.GetAllNendoroids(lang, character))
}

func (h *NendoroidHandler) GetNendoroidById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	// Handle languages not supported. Currently defaults to english but may reconsider adding an error code.
	lang := c.Query("language")
	if !h.isSupportedLang(lang) {
		lang = "en"
	}

	nendo, err := h.repo.GetNendoroidById(id, lang)
	if err != nil {
		c.JSON(404, gin.H{"message": fmt.Sprintf("Nendoroid #%d not found.", id)})
	} else {
		c.IndentedJSON(http.StatusOK, nendo)
	}
}

func (h *NendoroidHandler) isSupportedLang(lang string) bool {
	for _, supported := range h.supportedLangs {
		if supported == lang {
			return true
		}
	}
	return false
}
