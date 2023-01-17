package userController

import (
	"book-api/model"
	repo "book-api/repository"
	"net/http"

	"github.com/gin-contrib/sessions"
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

	cookie, _ := c.Cookie("session_token")
	session := sessions.Default(c)
	session.Set("session_token", cookie)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "login success"})

}

func Logout(c *gin.Context) {

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
