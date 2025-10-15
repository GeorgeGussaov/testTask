package handlers

import (
	"errors"
	"log"
	"net/http"
	"task/models"
	"task/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			log.Println("null id")
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is null"})
			return
		}

		var input models.Subscription
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Println("invalid JSON body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.repo.Update(c.Request.Context(), id, &input)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
				return
			}
			log.Println("failed update sub in DB:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "succeded"})
	}
}
