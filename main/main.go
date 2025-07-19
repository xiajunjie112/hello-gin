package main

import "github.com/gin-gonic/gin"

// 版本依赖入门案例，运行成功后访问http://localhost:8080/hello
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello GoModules!",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
