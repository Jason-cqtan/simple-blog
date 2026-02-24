package handlers

import (
	"net/http"
	"strconv"

	"github.com/Jason-cqtan/simple-blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostHandler struct {
	db *gorm.DB
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{db: db}
}

func (h *PostHandler) Home(c *gin.Context) {
	var posts []models.Post
	h.db.Preload("Author").Where("published = ?", true).Order("created_at desc").Limit(10).Find(&posts)
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home",
		"posts": posts,
	})
}

func (h *PostHandler) List(c *gin.Context) {
	var posts []models.Post
	h.db.Preload("Author").Where("published = ?", true).Order("created_at desc").Find(&posts)
	c.HTML(http.StatusOK, "posts/list.html", gin.H{
		"title": "All Posts",
		"posts": posts,
	})
}

func (h *PostHandler) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "posts/detail.html", gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := h.db.Preload("Author").First(&post, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "posts/detail.html", gin.H{"error": "Post not found"})
		return
	}

	var comments []models.Comment
	h.db.Preload("Author").Where("post_id = ?", id).Find(&comments)

	c.HTML(http.StatusOK, "posts/detail.html", gin.H{
		"title":    post.Title,
		"post":     post,
		"comments": comments,
	})
}

func (h *PostHandler) ShowCreateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/create.html", gin.H{"title": "Create Post"})
}

func (h *PostHandler) Create(c *gin.Context) {
	userID, _ := c.Get("userID")

	title := c.PostForm("title")
	content := c.PostForm("content")
	excerpt := c.PostForm("excerpt")
	category := c.PostForm("category")
	tags := c.PostForm("tags")

	if title == "" {
		c.HTML(http.StatusBadRequest, "posts/create.html", gin.H{"error": "Title is required"})
		return
	}

	post := models.Post{
		Title:     title,
		Content:   content,
		Excerpt:   excerpt,
		Category:  category,
		Tags:      tags,
		AuthorID:  userID.(uint),
		Published: true,
	}

	if err := h.db.Create(&post).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "posts/create.html", gin.H{"error": "Failed to create post: " + err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/posts/"+strconv.Itoa(int(post.ID)))
}

func (h *PostHandler) ShowEditForm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "posts/edit.html", gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := h.db.First(&post, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "posts/edit.html", gin.H{"error": "Post not found"})
		return
	}

	userID, _ := c.Get("userID")
	if post.AuthorID != userID.(uint) {
		c.HTML(http.StatusForbidden, "posts/edit.html", gin.H{"error": "Forbidden"})
		return
	}

	c.HTML(http.StatusOK, "posts/edit.html", gin.H{
		"title": "Edit Post",
		"post":  post,
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := h.db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	userID, _ := c.Get("userID")
	if post.AuthorID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	post.Title = c.PostForm("title")
	post.Content = c.PostForm("content")
	post.Excerpt = c.PostForm("excerpt")
	post.Category = c.PostForm("category")
	post.Tags = c.PostForm("tags")

	if err := h.db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post: " + err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/posts/"+strconv.Itoa(id))
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := h.db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	userID, _ := c.Get("userID")
	if post.AuthorID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if err := h.db.Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post: " + err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/posts")
}
