package api

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
	"github.com/mickaelyoshua/personal_finances/views"
)

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
func HandleRenderError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template - " + err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password - " + err.Error()})
		return
	}

	// Create a new user
	_, err = server.agent.CreateUser(c.Request.Context(), sqlc.CreateUserParams{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			switch pgErr.Code {
			case "23505": // Unique violation
				c.JSON(http.StatusForbidden, gin.H{"error": "Email already exists"})
				return
			}
			log.Println("PostgreSQL error:", pgErr.Code)
			log.Println("PostgreSQL error:", pgErr.Message)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user - " + err.Error()})
		return
	}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Compare the provided password with the stored hashed password
	err = util.CompareHashPassword(user.PasswordHash, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password - " + err.Error()})
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token - " + err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, int(server.config.AccessTokenDuration.Seconds()), "/", "", false, true)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/")
}

//*******************************************************Index Handler*******************************************************//
func (server *Server) Index(c *gin.Context) {
	err := Render(c, http.StatusOK, views.Index())
	HandleRenderError(c, err)
}