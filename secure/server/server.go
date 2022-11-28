package main

import (
	"example2.com/api"
	"example2.com/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Main()

	r := gin.Default()
	r.NoRoute(api.FOF)

	//Static Routes
	r.GET("/", api.Home)
	r.GET("/login", api.Login)
	r.POST("/login", api.LoginUser)
	r.GET("/testimonies", api.Testimonies)
	r.GET("/survey", api.Survey)
	r.GET("/gallery", api.Gallery)
	r.GET("/contact-us", api.ContactUs)
	r.POST("/contact-us", api.SendRequest)
	r.GET("/admin", api.Admin)

	//Load HTML Templates
	r.Static("/assets", "../website/assets")
	r.LoadHTMLGlob("../website/*.html")

	//specify ports
	r.Run(":8080")
}
