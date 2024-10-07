package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	//false let no. of capital
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue train", Artist: "Sid", Price: 100.1},
	{ID: "2", Title: "Pink train", Artist: "Ankit", Price: 200.1},
	{ID: "3", Title: "Red train", Artist: "Manish", Price: 305.1},
}

func getAlbums(c *gin.Context) { // gin context -> request details,validates,serializes
	c.IndentedJSON(http.StatusOK, albums)
	//return 200 response and data as converted albums to json
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")

}

func postAlbums(c *gin.Context) { // gin context -> request details,validates,serializes
	// steps

	// 1. get the data in json
	// 2.convert the data in our struct form
	var newAlbum album
	err := c.BindJSON(&newAlbum) //convert to struct and then assign to newAlbum

	// 2.1 report error is not valid json
	if err != nil {
		return
	}

	//3 update the data acc
	albums = append(albums, newAlbum)

	// 4. convey success to the client with the new data added
	//return 200 response and data as converted albums to json
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(c *gin.Context) { // gin context -> request details,validates,serializes
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album with the request id is not found"})
}
