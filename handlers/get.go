package handlers

import (
	"errors"
	"log"
	"net/http"
	"task/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		subscription, err := h.repo.SelectById(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
				return
			}
			log.Printf("error get subscribtion: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, subscription)
	}
}
