package handler

import (
	"test/internal/service/storage"
	"test/internal/http-server/middleware/logger"
	"github.com/gin-gonic/gin"
	// "test/internal/models/save"
	// "github.com/go-playground/validator/v10"
	// "net/http"
)


type Handler struct{
	service *storage.Service
}

func NewHandler(service *storage.Service) *Handler{
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







// //



func (h *Handler) ShortnerRoute(router *gin.Engine){
	api := router.Group("/shortner")
	{
		api.GET("/:id",h.GetURLbyId)
		api.POST("",h.CreateURL)
	}
}

func (h *Handler) InitRouter() *gin.Engine{
	router := gin.New()
	router.Run(":8080")
	router.Use(logger.NewMiddleware())
	// h.ShortnerRoute(router)
	return router
	
}

func New() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}