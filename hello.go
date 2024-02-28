package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Musics []string `json:"musics"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Dialogos", Artist: "Bardek", Musics: []string{"Dialogos", "Agora", "Amanhã", "Sempre"}, Price: 19.99},
}

var defaultAlbumValues = map[string]interface{}{
	"id":     "não achamos nos nossos servidores",
	"title":  "tem algo de errado aqui",
	"artist": "não foi cadastrado artista para este album",
	"musics": []string{"não temos nenhuma musica neste album ainda"},
	"price":  0.0,
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func setDefaultIfEmpty(newAlbum *album, defaults map[string]interface{}) {
	for key, value := range defaults {
		switch key {
		case "id":
			if newAlbum.ID == "" {
				newAlbum.ID = value.(string)
			}
		case "title":
			if newAlbum.Title == "" {
				newAlbum.Title = value.(string)
			}
		case "artist":
			if newAlbum.Artist == "" {
				newAlbum.Artist = value.(string)
			}
		case "musics":
			if newAlbum.Musics == nil {
				newAlbum.Musics = value.([]string)
			}
		case "price":
			if newAlbum.Price == 0 {
				newAlbum.Price = value.(float64)
			}
		}
	}
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context){
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}
	setDefaultIfEmpty(&newAlbum, defaultAlbumValues)

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	for _, a := range albums{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
