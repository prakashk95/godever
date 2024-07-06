package main

import (
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func blog(c *gin.Context) {
	c.HTML(200, "blog.html", nil)
}
func learn(c *gin.Context) {
	c.HTML(200, "learn.html", nil)
}
func helloworld(c *gin.Context) {
	c.HTML(200, "hello-world.html", nil)
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", home)
	r.GET("/hello-world", helloworld)
	r.GET("/blog", blog)
	r.GET("/learn", learn)

	r.Run("localhost:1200")
}
