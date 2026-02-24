package handlers

import (
	"net/http"
	"strconv"

	"github.com/Jason-cqtan/simple-blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandler struct {
	db *gorm.DB
}

func NewCommentHandler(db *gorm.DB) *CommentHandler {
	return &CommentHandler{db: db}
}

func (h *CommentHandler) Create(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	userID, _ := c.Get("userID")
	content := c.PostForm("content")

	if content == "" {
		c.Redirect(http.StatusFound, "/posts/"+strconv.Itoa(postID))
		return
	}

	comment := models.Comment{
		Content:  content,
		AuthorID: userID.(uint),
		PostID:   uint(postID),
	}

	if err := h.db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.Redirect(http.StatusFound, "/posts/"+strconv.Itoa(postID))
}
