package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

    r.Run(":8080") // Menjalankan aplikasi di port 8080
}
