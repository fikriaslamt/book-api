package bookController

import (
	model "book-api/model"
	repo "book-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	data, err := repo.GetAllBooks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "InternalServerError",
			"error":  err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"book": data})
}

func GetBook(c *gin.Context) {
	idParam := c.Param("id")
	idBook, _ := strconv.Atoi(idParam)

	data, err := repo.GetBook(idBook)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "InternalServerError",
			"error":  err.Error(),
		})
		return

	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"book": data})
}

func AddBook(c *gin.Context) {
	var book = model.Book{}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "BadRequest",
			"error":  err.Error(),
		})
		return
	}

	if err := repo.AddBook(book); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "InternalServerError",
			"error":  err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"book": book})

}

func UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	idBook, _ := strconv.Atoi(idParam)
	var book = model.Book{}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "BadRequest",
			"error":  err.Error(),
		})
		return
	}

	if err := repo.UpdateBook(idBook, book); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "InternalServerError",
			"error":  err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"book": book})

}

func DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	idBook, _ := strconv.Atoi(idParam)

	if err := repo.DeleteBook(idBook); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "InternalServerError",
			"error":  err.Error(),
		})
		return

	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, nil)
}
