package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	OTP      string `form:"otp"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if err := c.Bind(&form); err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": err.Error(),
			})
			return
		}

		// Perform necessary validations and actions based on the email/phone number and OTP/password
		if form.Email != "" {
			// Handle email login
			if form.Password == "" {
				c.HTML(http.StatusBadRequest, "index.html", gin.H{
					"error": "Password is required for email login",
				})
				return
			}
			// Call external services or databases for email login
		} else {
			// Handle phone number login
			if form.OTP == "" {
				c.HTML(http.StatusBadRequest, "index.html", gin.H{
					"error": "OTP is required for phone number login",
				})
				return
			}
			// Call external services or databases for phone number login
		}

		// Handle successful login
		c.Redirect(http.StatusFound, "/success")
	})

	r.GET("/success", func(c *gin.Context) {
		c.HTML(http.StatusOK, "success.html", nil)
	})

	r.Run()
}
