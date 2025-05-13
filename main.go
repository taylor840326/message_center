package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		println(string(body))
		c.JSON(200, "ok")
	})
	r.Run(":9777") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
