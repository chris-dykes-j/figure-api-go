package main

import (
	h "figures/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler := h.Init()

	router.GET("/", handler.GetHomePage)
	router.GET("/nendoroid", handler.GetAllNendoroids)
	router.GET("/nendoroid/:id", handler.GetNendoroidById)

	router.Run("localhost:8080")
}
