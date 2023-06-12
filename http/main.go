package main

import (
	"bytes"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/mendoncas/gostify/service"
)

func main() {
	router := gin.Default()

	router.GET("/users", service.RestGetUsers)
	router.GET("/playlists", service.RestGetPlaylists)
	router.GET("/songs", service.RestGetSongs)
	router.POST("/soap", handleSoapRequest)
	router.Run("localhost:8085")
}

type SoapRequest struct {
	Action string `xml:"action"`
}

func handleSoapRequest(c *gin.Context) {
	var x SoapRequest
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	xml.Unmarshal(buf.Bytes(), &x)

	switch x.Action {
	case "getUsers":
		service.SOAPGetUsers(c)
	case "getPlaylists":
		service.SOAPGetPlaylists(c)
	case "getSongs":
		service.SOAPGetSongs(c)
	default:
		c.XML(400, "requisição inválida")
	}
}
