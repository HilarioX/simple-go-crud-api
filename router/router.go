package router

import (
	"rest-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/abums", controller.CreateAlbum)
	router.Run("localhost:3333")
}
