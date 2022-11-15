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
	r.GET("/testimonies", api.Testimonies)
	r.GET("/survey", api.Survey)
	r.GET("/gallery", api.Gallery)
	r.GET("/contact-us", api.ContactUs)
	r.GET("/admin", api.Admin)

	//Load HTML Templates
	r.Static("/assets", "../website/assets")
	r.LoadHTMLGlob("../website/*.html")

	r.Run()
}
