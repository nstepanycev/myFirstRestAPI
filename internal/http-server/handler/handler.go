package handler

import (
	"test/internal/service/storage"
	"test/internal/http-server/middleware/logger"
	"github.com/gin-gonic/gin"
)


type Handler struct{
	service *storage.Service
}

func NewHandler(service *storage.Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine{
	router := gin.New()
	router.Use(logger.NewMiddleware())

	h.ShortnerRoute(router)

	return router
}
