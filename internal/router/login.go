package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mazlon/gobeyond/internal/auth"
	"github.com/mazlon/gobeyond/internal/models"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (gbt GbtServer) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer c.Request.Body.Close()
	userId, err := models.UserExists(gbt.dbConnection, input.Username, input.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.GenerateToken(userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = SetCookie(c, Cookies{"token": token, "user_id": userId})
	if err != nil {
		log.Printf("an Error while setting token as cookie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured"})
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "authenticated"})

}
