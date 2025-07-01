package api

import (
	"log"
	"net/http"
	"regexp"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/token"
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
	err := Render(c, http.StatusOK, views.Register(sqlc.User{}, nil))
	HandleRenderError(c, err)
}

func LoginView(c *gin.Context) {
	err := Render(c, http.StatusOK, views.Login())
	HandleRenderError(c, err)
}

func validateRegisterForm(c *gin.Context, server *Server, name, email, password, confirmPassword string) map[string]string {
	errors := make(map[string]string, 4)

	// Validate name: only alphabetic and between 3 and 50 characters
	if len(name) < 3 || len(name) > 50 {
		errors["name"] = "Name must be 3-50 characters"
	} else if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
		errors["name"] = "Name must contain letters only"
	}

	// Validate email: must have a valid email format
	if !regexp.MustCompile(util.EmailRegexPattern).MatchString(email) {
		errors["email"] = "Invalid email format"
	}

	// Check if email already exists
	user, err := server.Agent.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			// Email is available
		} else {
			errors["email"] = "Check email got unexpected error: " + err.Error()
		}
	}
	if user.ID != 0 || user.Email == email {
		errors["email"] = "Email already exists"
	}

	// Validate password: must be at least 6 characters long
	if len(password) < 6 {
		errors["password"] = "Password must be at least 6 characters long"
	}

	// Validate confirm password: must match password
	if password != confirmPassword {
		errors["confirm_password"] = "Passwords do not match"
	}

	return errors
}
func (server *Server) Register(c *gin.Context) {
	// Get form values
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	// Validate form
	errors := validateRegisterForm(c, server, name, email, password, confirmPassword)
	if len(errors) > 0 {
		log.Println("Validation errors:", errors)
		err := Render(c, http.StatusBadRequest, views.RegisterForm(sqlc.User{Name: name, Email: email}, errors))
		HandleRenderError(c, err)
		return
	}

	// Hash the password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password - " + err.Error()})
		return
	}

	// Create a new user
	user, err := server.Agent.CreateUser(c.Request.Context(), sqlc.CreateUserParams{
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
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error - " + pgErr.Message + " (code: " + pgErr.Code + ")"})
				log.Println("PostgreSQL error:", pgErr.Code)
				log.Println("PostgreSQL error:", pgErr.Message)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user - " + err.Error()})
		return
	}

	server.SetToken(c, user.ID)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/")
}

func (server *Server) Login(c *gin.Context) {
	// Get form values
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Get user by email
	user, err := server.Agent.GetUserByEmail(c.Request.Context(), email)
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

	server.SetToken(c, user.ID)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/")
}

//*******************************************************Index Handler*******************************************************//
func (server *Server) Index(c *gin.Context) {
	payload, exists := c.Get(authorizationPayloadKey)
	if !exists {
		log.Println("Authorization payload not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization payload not found"})
		return
	}

	authPayload := payload.(*token.Payload)
	token, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found in cookies"})
		return
	}

	tokenPayload, err := server.TokenMaker.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token - " + err.Error()})
		return
	}

	if tokenPayload.UserID != authPayload.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token does not match user ID in context"})
		return
	}

	user, err := server.Agent.GetUserById(c.Request.Context(), authPayload.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user - " + err.Error()})
		return
	}

	err = Render(c, http.StatusOK, views.Index(user))
	HandleRenderError(c, err)
}