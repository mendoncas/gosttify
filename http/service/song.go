package service

import "github.com/gin-gonic/gin"

type song struct {
	Title    string  `json:"title" xml:"title"`
	Duration float64 `json:"duration" xml:"duration"`
	Artist   string  `json:"artist" xml:"artist"`
}

var songs = [6]song{
	{Title: "Frances farmer will have her revenge", Duration: 4.2, Artist: "Nirvana"},
	{Title: "Down in a hole", Duration: 4.2, Artist: "Alice in Chains"},
	{Title: "Doomsday", Duration: 4.2, Artist: "MF DOOM"},
	{Title: "Fullgás", Duration: 4.2, Artist: "Marina Lima"},
	{Title: "Viking hair", Duration: 4.2, Artist: "Dry Cleaning"},
	{Title: "Get Got", Duration: 4.2, Artist: "Death Grips"},
}

func RestGetSongs(c *gin.Context) {
	c.IndentedJSON(200, songs)
}

func SOAPGetSongs(c *gin.Context) {
	c.XML(200, songs)
}
