package server

import (
	"log"
	"net/http"
	"strconv"

	"csuProject/storage"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Provider *Provider
	Repo     *storage.Repository
}

func (s *Router) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/hello", hello)
	links := router.Group("/link")
	{
		links.GET("/:id", s.getLink)
		links.POST("/:link", s.CreateLink)
		links.PUT("/:id/:status", s.UpdateLink)
	}
	return router
}

func (s *Router) getLink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_reason": "неверные параметры запроса",
		})
		return
	}
	link, err := s.Repo.GetLink(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"link": link,
	})
}

func (s *Router) CreateLink(c *gin.Context) {
	link := c.Param("link")
	id, err := s.Repo.CreateLink(link)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	s.Provider.Publish(id, link)
	c.JSON(http.StatusOK, map[string]interface{}{
		"new_link_id": id,
	})
}

func hello(c *gin.Context) {
	c.Writer.Write([]byte("<h1>hello world!!<h1>"))
}

func (s *Router) UpdateLink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_reason": "неверные параметры запроса: id",
		})
		return
	}
	status := c.Param("status")
	log.Println("updating url")
	s.Repo.UpdateLink(id, status)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}
