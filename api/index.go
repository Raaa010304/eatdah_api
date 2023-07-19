package handler

import (
	"./controllers/menucontroller"
	"./controllers/menufavoritecontroller"
	"./models"

	"github.com/gin-gonic/gin"
	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gin.Default()

	models.ConnectDatabase()

	server.GET("/api/menus", menucontroller.Index)
	server.GET("/api/menu/:id", menucontroller.Show)
	server.POST("/api/menu", menucontroller.Create)
	server.PUT("/api/menu/:id", menucontroller.Update)
	server.DELETE("/api/menu", menucontroller.Delete)

	server.GET("/api/menusfavorite", menufavoritecontroller.Index)
	server.GET("/api/menufavorite/:id", menufavoritecontroller.Show)
	server.POST("/api/menufavorite", menufavoritecontroller.Create)
	server.PUT("/api/menufavorite/:id", menufavoritecontroller.Update)
	server.DELETE("/api/menufavorite", menufavoritecontroller.Delete)

	server.Handle(w, r)
}
