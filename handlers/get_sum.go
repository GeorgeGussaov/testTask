package handlers

import (
	"log"
	"net/http"
	"task/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSum() gin.HandlerFunc {
	return func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")
		userID := c.Query("user_id")
		name := c.Query("service_name")

		var info models.SubscriptionFilter

		if from == "" || to == "" { //как я понял период указывается обязательно
			c.JSON(http.StatusBadRequest, gin.H{"error": "specify the period"})
			return
		}
		f, err := time.Parse("2006-01-02", from)
		t, err := time.Parse("2006-01-02", to)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "specify the valid period (format yyyy-mm-dd)"}) //надеюсь с форматом не принципиально
			return
		}
		info.FromDate = &f
		info.ToDate = &t

		if f.After(t) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "from must be <= to"})
			return
		}
		if userID != "" {
			info.UserID = &userID
		}
		if name != "" {
			info.ServiceName = &name
		}

		sum, err := h.repo.SelectSumByInfo(c.Request.Context(), &info)
		if err != nil {
			log.Println("DB error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get subscriptions"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sum": sum})
	}
}
