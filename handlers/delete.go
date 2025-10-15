package handlers

import (
	"errors"
	"log"
	"net/http"
	"task/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := h.repo.DeleteById(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
				return
			}
			log.Printf("error delete subscribtion: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
