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

func Testimonies(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"testimonies",
		gin.H{},
	)
}

func Gallery(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"gallery",
		gin.H{},
	)
}

func Survey(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"survey",
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
