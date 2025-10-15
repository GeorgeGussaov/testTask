package handlers

import (
	"task/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *repository.SubscriptionsRepository
}

func NewHandler(r *repository.SubscriptionsRepository) *Handler {
	return &Handler{repo: r}
}

func (h *Handler) RegisterRoutes(rg *gin.Engine) {
	r := rg.Group("/subscriptions")
	r.POST("", h.Post())
	r.PUT("/:id", h.Put())
	r.GET("/:id", h.Get())
	r.GET("", h.GetList())
	r.GET("/sum", h.GetSum())
	r.DELETE("/:id", h.Delete())
}
