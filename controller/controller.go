package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"rest-go/models"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func CreateAlbum(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var newAlbum models.Album
	err = json.Unmarshal(body, &newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON"})
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
