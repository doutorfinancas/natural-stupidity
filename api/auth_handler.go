package api

import (
	"net/http"

	"github.com/doutorfinancas/natural-stupidity/service"
	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes registers login and logout endpoints.
func RegisterAuthRoutes(r *gin.Engine, authSvc service.AuthService) {
	r.GET("/login", showLoginPage)
	r.POST("/login", loginHandler(authSvc))
	r.GET("/logout", logoutHandler)
}

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

type loginForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func loginHandler(authSvc service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form loginForm
		if err := c.ShouldBind(&form); err != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": err.Error()})
			return
		}
		user, err := authSvc.Authenticate(form.Email, form.Password)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid credentials"})
			return
		}
		// Generate JWT
		token, err := authSvc.GenerateToken(user)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": err.Error()})
			return
		}
		// Set cookie
		c.SetCookie("jwt", token, 3600*24, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
	}
}

func logoutHandler(c *gin.Context) {
	// Clear cookie
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}
