package api

import (
	"fmt"
	"net/http"

	"example.com/db"
	"example.com/query"
	"github.com/gin-gonic/gin"
)

var conn = db.Main()

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
		gin.H{
			"status": "none",
		},
	)
}

func LoginUser(c *gin.Context) {
	c.Request.ParseForm()
	// user, ok1 := c.Request.PostForm["user"]
	// password, ok2 := c.Request.PostForm["password"]

	username := c.Request.PostForm["user"]
	password := c.Request.PostForm["password"]
	fmt.Printf("username: %s", username)
	fmt.Printf("password: %s", password)
	val := query.NewUser(conn, username[0], password[0])
	c.HTML(
		http.StatusOK,
		"login",
		gin.H{
			"status": val,
		},
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

func SendRequest(c *gin.Context) {
	name := c.Request.PostFormValue("name")
	email := c.Request.PostFormValue("email")
	desc := c.Request.PostFormValue("desc")
	fmt.Printf("name: %s", name)
	fmt.Printf("email: %s", email)
	fmt.Printf("desc: %s", desc)
	status := query.NewRequest(conn, name, email, desc)
	c.HTML(
		http.StatusOK,
		"contact-us",
		gin.H{
			"status": status,
		},
	)
}

func Admin(c *gin.Context) {
	users := query.QueryUsers(conn)
	requests := query.QueryRequests(conn)
	c.HTML(
		http.StatusOK,
		"admin",
		gin.H{
			"Users":    users,
			"Requests": requests,
		},
	)
}

func FOF(c *gin.Context) {
	c.HTML(
		http.StatusNotFound,
		"fof",
		gin.H{},
	)
}
