package router

import (
	"github.com/mazlon/gobeyond/internal/handlers"
	"github.com/gin-gonic/gin"
)

func PrepareRoutes(r *gin.Engine) {
	r.GET("/hello", handlers.HelloWorld)
	r.POST("/ask", handlers.AskQuestions)
}
