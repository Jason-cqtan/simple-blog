package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Jason-cqtan/simple-blog/models"
	"gorm.io/gorm"
)

// Seed populates the database with default admin user, sample posts, and
// sample comments.  All operations are idempotent – running Seed multiple
// times will not create duplicate records.
func Seed(db *gorm.DB) error {
	admin, err := seedAdminUser(db)
	if err != nil {
		return fmt.Errorf("seed admin user: %w", err)
	}

	posts, err := seedPosts(db, admin.ID)
	if err != nil {
		return fmt.Errorf("seed posts: %w", err)
	}

	if err := seedComments(db, admin.ID, posts); err != nil {
		return fmt.Errorf("seed comments: %w", err)
	}

	log.Println("Database seeding completed successfully.")
	return nil
}

// seedAdminUser creates the default admin account if it does not exist yet.
func seedAdminUser(db *gorm.DB) (*models.User, error) {
	admin := &models.User{}
	result := db.Where("username = ?", "admin").First(admin)
	if result.Error == nil {
		log.Println("Admin user already exists, skipping.")
		return admin, nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	admin = &models.User{
		Username: "admin",
		Email:    "admin@example.com",
		Bio:      "Default administrator account.",
	}
	if err := admin.HashPassword("Admin@123456"); err != nil {
		return nil, fmt.Errorf("hash admin password: %w", err)
	}
	if err := db.Create(admin).Error; err != nil {
		return nil, err
	}
	log.Printf("Created admin user (id=%d).", admin.ID)
	return admin, nil
}

// seedPosts creates sample blog posts authored by adminID if none exist yet.
func seedPosts(db *gorm.DB, adminID uint) ([]models.Post, error) {
	var count int64
	db.Model(&models.Post{}).Count(&count)
	if count > 0 {
		log.Println("Posts already exist, skipping.")
		var existing []models.Post
		db.Find(&existing)
		return existing, nil
	}

	samples := []models.Post{
		{
			Title:     "Welcome to Simple Blog",
			Content:   "This is the first post on Simple Blog. Feel free to explore the features: create posts, leave comments, and manage your profile.",
			Excerpt:   "Welcome to Simple Blog – your new Go-powered blogging platform.",
			AuthorID:  adminID,
			Category:  "General",
			Tags:      "welcome,intro",
			Published: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Getting Started with Go and Gin",
			Content:   "Go is a statically typed, compiled language designed for simplicity and performance. The Gin framework makes it easy to build fast HTTP servers with clean routing and middleware support.",
			Excerpt:   "A brief introduction to building web applications with Go and the Gin framework.",
			AuthorID:  adminID,
			Category:  "Technology",
			Tags:      "go,gin,web",
			Published: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Using PostgreSQL with GORM",
			Content:   "GORM is a full-featured ORM library for Go. Combined with PostgreSQL it provides powerful tools for schema migration, associations, and querying.",
			Excerpt:   "Learn how to use GORM with PostgreSQL in a Go web application.",
			AuthorID:  adminID,
			Category:  "Technology",
			Tags:      "go,gorm,postgresql",
			Published: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(&samples).Error; err != nil {
		return nil, err
	}
	log.Printf("Created %d sample posts.", len(samples))
	return samples, nil
}

// seedComments creates sample comments on the provided posts if none exist yet.
func seedComments(db *gorm.DB, adminID uint, posts []models.Post) error {
	var count int64
	db.Model(&models.Comment{}).Count(&count)
	if count > 0 {
		log.Println("Comments already exist, skipping.")
		return nil
	}
	if len(posts) == 0 {
		return nil
	}

	comments := []models.Comment{
		{
			Content:   "Great first post! Looking forward to more content.",
			AuthorID:  adminID,
			PostID:    posts[0].ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Content:   "Gin is indeed an excellent framework. Thanks for sharing!",
			AuthorID:  adminID,
			PostID:    posts[1].ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Content:   "GORM's AutoMigrate feature is a real time-saver.",
			AuthorID:  adminID,
			PostID:    posts[2].ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(&comments).Error; err != nil {
		return err
	}
	log.Printf("Created %d sample comments.", len(comments))
	return nil
}
