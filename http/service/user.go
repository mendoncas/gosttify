package service

import "github.com/gin-gonic/gin"

type User struct {
	Username  string     `json:"username" xml:"username"`
	Playlists []playlist `json:"playlists" xml:"playlists"`
}

var playlistsRicardo = []playlist{playlists[0]}
var playlistsPedro = []playlist{playlists[1]}

var users = [2]User{
	{Username: "ricardo", Playlists: playlistsRicardo},
	{Username: "pedro", Playlists: playlistsPedro},
}

func RestGetUsers(c *gin.Context) {
	c.IndentedJSON(200, users)
}

func SOAPGetUsers(c *gin.Context) {
	c.XML(200, users)
}

func GetUsers() []User {
	return users[:]
}
