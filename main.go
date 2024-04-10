package main

import (
	h "figures/handlers"
	r "figures/repositories"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

    conn_str := os.Getenv("FIGURE_DB")
    repo := r.NewNendoRepository(conn_str)
	handler := h.NewNendoHandler(repo)

	router.GET("/", handler.GetHomePage)
	router.GET("/nendoroid", handler.GetAllNendoroids)
	router.GET("/nendoroid/:id", handler.GetNendoroidById)

	router.Run("localhost:8080")
}
