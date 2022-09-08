package server

import "github.com/gin-gonic/gin"

type Router struct {
}

func (s *Router) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/hello", hello)
	return router
}

func hello(c *gin.Context) {
	c.Writer.Write([]byte("<h1>hello world!!<h1>"))
}
