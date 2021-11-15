package middleware

import "github.com/gin-gonic/gin"

type authenticate struct{}

func NewAuth() *authenticate {
	return &authenticate{}
}

func (a *authenticate) Authenticate(c *gin.Context) {
	
}