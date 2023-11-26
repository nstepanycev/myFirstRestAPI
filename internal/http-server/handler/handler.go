package handler

import (
	"test/internal/http-server/middleware"
	"test/internal/service/storage"

	"github.com/gin-gonic/gin"
)


type Handler struct{
	service *service.storage
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) initRouter() *gin.Engine{
	router := gin.New()
	router.Use(middleware.Error())

	h.ShortnerRoute(router)

	return router
}
