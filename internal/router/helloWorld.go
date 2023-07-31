package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mazlon/gobeyond/internal/auth"
	"github.com/mazlon/gobeyond/internal/models"
)

func (gbt *GbtServer) HelloWorld(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Printf("An error while extracting the token cookie: %v", err)
		c.JSON(401, gin.H{
			"Error": "Authentication required",
		})
		return
	}
	// userID, err := c.Cookie("user_id")
	// if err != nil {
	// 	log.Printf("An error while extracting the token cookie: %v", err)
	// 	c.JSON(500, gin.H{
	// 		"Error": "An error with the cookies",
	// 		"Message": "You can not get authenticated",
	// 	})
	// 	return
	// }
	expected_uID, err := auth.TokenVerifier(token)
	if err != nil {
		log.Printf("An error while extracting the token cookie: %v", err)
		c.JSON(403, gin.H{
			"Error": "You're unable to get authenticated",
		})
		return
	}
	username, err := models.GetUserName(gbt.dbConnection, expected_uID)
	if err != nil || username == "" {
		log.Printf("couldn't find user for the userID %s error: %v", expected_uID, err)
		c.JSON(500, gin.H{
			"Error": "An error with finding the username",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Hello, %s!", username),
	})
}
