package handlers

import (
	"fmt"
	"log"
	"net/http"
	"task/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Subscription
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Println("Invalid JSON body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := h.repo.Insert(c.Request.Context(), input)
		if err != nil {
			log.Println("failed insert sub in DB:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Header("Location", fmt.Sprintf("/subscriptions/%s", id))
		c.JSON(http.StatusCreated, gin.H{
			"status": "succeeded",
			"id":     id,
		})
	}
}
