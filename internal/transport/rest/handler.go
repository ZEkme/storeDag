package rest

import (
	"github.com/Way-Flare/dagestan-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{services: &services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		api.GET("/place", h.getAllPlaces)
		api.GET("/place/id", h.getPlaceById)
	}

	return router
}
