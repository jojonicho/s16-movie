package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	movie "github.com/jojonicho/s16-movie/movie"
)

type MovieHTTPHandler interface {
	Search(*gin.Context)
	Detail(*gin.Context)
}

type movieHTTPHandler struct {
	MovieSvc movie.Service
}

func NewMovieHTTPHandler(movieSvc movie.Service) MovieHTTPHandler {
	return &movieHTTPHandler{
		MovieSvc: movieSvc,
	}
}

func (h *movieHTTPHandler) Search(c *gin.Context) {
	data, err := h.MovieSvc.Search(c, movie.SearchOpts{
		Title: c.Query("s"),
		Type:  c.Query("type"),
		Year:  c.Query("y"),
		Page:  c.Query("page"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *movieHTTPHandler) Detail(c *gin.Context) {
	data, err := h.MovieSvc.Detail(c, movie.DetailOpts{
		ID:    c.Query("i"),
		Title: c.Query("t"),
		Type:  c.Query("type"),
		Year:  c.Query("y"),
		Plot:  c.Query("plot"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
