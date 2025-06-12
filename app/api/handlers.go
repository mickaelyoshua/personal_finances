package api

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal-finances/db/sqlc"
	"github.com/mickaelyoshua/personal-finances/util"
	"github.com/mickaelyoshua/personal-finances/views"
)

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
func HandleRenderError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		return
	}
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

//*******************************************************Auth Handlers*******************************************************//
func RegisterView(c *gin.Context) {
	err := Render(c, http.StatusOK, views.Register())
	HandleRenderError(c, err)
}

func LoginView(c *gin.Context) {
	err := Render(c, http.StatusOK, views.Login())
	HandleRenderError(c, err)
}

func (server *Server) Register(c *gin.Context) {
	// Get form values
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Hash the password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user
	user, err := server.agent.CreateUser(c.Request.Context(), sqlc.CreateUserParams{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate a token for the user
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Set the token as a cookie
	c.SetCookie("token", token, int(72*time.Hour.Seconds()), "/", "", false, true)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/")
}

func (server *Server) Login(c *gin.Context) {
	// Get form values
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Get user by email
	user, err := server.agent.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// Compare the provided password with the stored hashed password
	if !util.CompareHashedPassword(user.PasswordHash, password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate a token for the user
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Set the token as a cookie
	c.SetCookie("token", token, int(72*time.Hour.Seconds()), "/", "", false, true)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/")
}

//*******************************************************Index Handler*******************************************************//
func (server *Server) Index(c *gin.Context) {
	// Get token from cookie
	token, err := util.GetTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - " + err.Error()})
		return
	}

	// Parse and validate the token
	claims, err := util.ParseAndValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - " + err.Error()})
		return
	}

	// Get user from userID
	userID := int32(claims["userID"].(float64))
	user, err := server.agent.GetUserById(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user from database - " + err.Error()})
		return
	}

	err = Render(c, http.StatusOK, views.Index(user))
	HandleRenderError(c, err)
}