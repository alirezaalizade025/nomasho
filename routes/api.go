package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nomasho/controllers/auth"
)

func Register()  {
	r := gin.Default()


	// ping
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})

	// authentication
	r.POST("/login", auth.Register)
		

	r.Run(":8181")
}