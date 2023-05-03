package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mazlon/gobeyond/internal/router"
)

func main() {
    r := gin.Default()
    router.PrepareRoutes(r)
    // Run the server
    r.Run(":8080")
}
