package middleware

import "github.com/gin-gonic/gin"


type MiddleWareErrors struct{
	Message string `json:"massage"`
	Code    string `json:"code"`
}

func Error() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()
	}
}