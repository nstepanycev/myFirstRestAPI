package handler

import (
	"github.com/gin-gonic/gin"

)

func (h *Handler) New() *gin.Engine{
	router := gin.New()

	router.Use(.RequestID())
	router.Use(middleware.Logger())
	router.Use(middleware.Recoverer())
	router.Use(middleware.URLFormat())

	h.UrlShortnerRouter(router)

}

func (h *Handler) NewRequest()