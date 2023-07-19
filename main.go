package main

import (
	"eatdah/controllers/menucontroller"
	"eatdah/controllers/menufavoritecontroller"
	"eatdah/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/menus", menucontroller.Index)
	r.GET("/api/menu/:id", menucontroller.Show)
	r.POST("/api/menu", menucontroller.Create)
	r.PUT("/api/menu/:id", menucontroller.Update)
	r.DELETE("/api/menu", menucontroller.Delete)

	r.GET("/api/menusfavorite", menufavoritecontroller.Index)
	r.GET("/api/menufavorite/:id", menufavoritecontroller.Show)
	r.POST("/api/menufavorite", menufavoritecontroller.Create)
	r.PUT("/api/menufavorite/:id", menufavoritecontroller.Update)
	r.DELETE("/api/menufavorite", menufavoritecontroller.Delete)

	r.Run()
}
