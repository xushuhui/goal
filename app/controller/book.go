package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func BookList(c *gin.Context) {
	claims, _ := c.Get("claims")
	fmt.Println(claims)
}
