package handlers

import (
	"net/http"
	"time"

	"github.com/Jason-cqtan/simple-blog/config"
	"github.com/Jason-cqtan/simple-blog/models"
	"github.com/Jason-cqtan/simple-blog/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewUserHandler(db *gorm.DB, cfg *config.Config) *UserHandler {
	return &UserHandler{db: db, cfg: cfg}
}

func (h *UserHandler) ShowRegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "users/register.html", gin.H{"title": "Register"})
}

func (h *UserHandler) Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if err := utils.ValidateUsername(username); err != nil {
		c.HTML(http.StatusBadRequest, "users/register.html", gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateEmail(email); err != nil {
		c.HTML(http.StatusBadRequest, "users/register.html", gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidatePassword(password); err != nil {
		c.HTML(http.StatusBadRequest, "users/register.html", gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: username,
		Email:    email,
	}
	if err := user.HashPassword(password); err != nil {
		c.HTML(http.StatusInternalServerError, "users/register.html", gin.H{"error": "Failed to process password"})
		return
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "users/register.html", gin.H{"error": "Username or email already exists"})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func (h *UserHandler) ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "users/login.html", gin.H{"title": "Login"})
}

func (h *UserHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	if err := h.db.Where("email = ?", email).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "users/login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	if !user.CheckPassword(password) {
		c.HTML(http.StatusUnauthorized, "users/login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID, h.cfg.JWTSecret)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "users/login.html", gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "", h.cfg.SecureCookie, true)
	c.Redirect(http.StatusFound, "/")
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/")
}

func (h *UserHandler) ShowProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.HTML(http.StatusNotFound, "users/profile.html", gin.H{"error": "User not found"})
		return
	}

	var posts []models.Post
	h.db.Where("author_id = ?", userID).Find(&posts)

	c.HTML(http.StatusOK, "users/profile.html", gin.H{
		"title": user.Username + "'s Profile",
		"user":  user,
		"posts": posts,
	})
}
