package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		subs, err := h.repo.SelectList(c.Request.Context())
		if err != nil {
			log.Println("failed get subs list:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get subscriptions"})
			return
		}
		c.JSON(http.StatusOK, subs)
	}
}
