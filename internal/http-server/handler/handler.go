package handler

import (
	"test/internal/http-server/middleware/logger"
	service "test/internal/services"
	// "test/internal/services/storage"

	"github.com/gin-gonic/gin"
	// "test/internal/models/save"
	// "github.com/go-playground/validator/v10"
	// "net/http"
)


type Handler struct{
	service *service.Service
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{service: service}
}

// Другой Хендлер

// type Storage interface{
// 	SaveURL(urlToSave *models.URLStorage) (int, error)
// 	// GetURL(id int) (*models.URLStorage, error)
// }

// func CreateURL(SaveURL Storage) gin.HandlerFunc{
// 	return func(c *gin.Context){
// 		var req models.Request

// 	}
// }


// func (h Handler) getServices(gin.Context){
// 	serviceListr, err := h.service.GetURL()
// }




// //



func (h *Handler) ShortnerRoute(router *gin.Engine){
	api := router.Group("/shortner")
	{
		api.GET("/:id",h.GetURLbyIdService)
		api.POST("",h.CreateURLService)
	}
}

func (h *Handler) InitRouter() *gin.Engine{
	router := gin.New()
	router.Run(":8080")
	router.Use(logger.NewMiddleware())
	// h.ShortnerRoute(router)
	return router
	
}