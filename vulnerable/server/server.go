package main

import (
	"example.com/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.NoRoute(api.FOF)

	//Static Routes
	r.GET("/", api.Home)
	r.GET("/login", api.Login)
	r.GET("/thing2", api.Thing2)
	r.GET("/thing1", api.Thing1)
	r.GET("/contact-us", api.ContactUs)
	r.GET("/admin", api.Admin)

	//Load HTML Templates
	r.Static("/assets", "../website/assets")
	r.LoadHTMLGlob("../website/*.html")

	r.Run()
}
