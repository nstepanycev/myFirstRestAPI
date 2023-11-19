package handler

import "github.com/gin-gonic/gin"


func (h *Handler) UrlShortnerRouter(router *gin.Engine){
	shortner := router.Group("/sh")
}