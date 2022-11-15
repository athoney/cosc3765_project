package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home",
		gin.H{},
	)
}

func Login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login",
		gin.H{},
	)
}

func Thing2(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"thing-2",
		gin.H{},
	)
}

func Thing1(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"thing-1",
		gin.H{},
	)
}

func ContactUs(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"contact-us",
		gin.H{},
	)
}

func Admin(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"admin",
		gin.H{},
	)
}

func FOF(c *gin.Context) {
	c.HTML(
		http.StatusNotFound,
		"fof",
		gin.H{},
	)
}
