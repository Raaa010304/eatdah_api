package handler

import (
	"eatdah/controllers/menucontroller"
	"eatdah/controllers/menufavoritecontroller"
	"eatdah/models"

	"github.com/gin-gonic/gin"
)

func Handler() {
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

	server.Run()
}
