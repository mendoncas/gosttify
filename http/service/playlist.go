package service

import "github.com/gin-gonic/gin"

type playlist struct {
	Title string `json:"title" xml:"title"`
	Songs []song `json:"songs" xml:"songs"`
}

var playlists = [2]playlist{
	{Title: "Playlist do Ricardo", Songs: songs[0:1]},
	{Title: "Playlist do Pedro", Songs: songs[2:]},
}

func RestGetPlaylists(c *gin.Context) {
	c.IndentedJSON(200, playlists)
}

func SOAPGetPlaylists(c *gin.Context) {
	c.XML(200, playlists)
}
