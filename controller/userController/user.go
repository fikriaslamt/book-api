package userController

import (
	"book-api/model"
	repo "book-api/repository"
	"book-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {

	user := model.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "BadRequest",
			"error":  err.Error(),
		})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "BadRequest",
			"error":  "Username or Password Empty",
		})
		return
	}

	err := repo.UserAvail(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "BadRequest",
				"error":  "Username or Password Not Found",
			})
			return
		}
	}

	tokenString := utils.GenerateToken(user.Username)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"auth",
		tokenString,
		3600*24,
		"",
		"",
		false,
		true,
	)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "login success"})

}

func Logout(c *gin.Context) {

	c.SetCookie(
		"auth",
		"",
		-1,
		"",
		"",
		false,
		true,
	)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
