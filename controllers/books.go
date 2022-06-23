package controllers

import (
	"example/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(ctx *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create new book
func CreateBook(ctx *gin.Context) {
	// Validate input
	var input models.CreateBookInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{}
	book.Title = input.Title
	book.Author = input.Author
	models.DB.Create((&book))

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(ctx *gin.Context) {
	// Check if the book exist
	var book models.Book
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate the update input
	var input models.UpdateBookInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(ctx *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	models.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
