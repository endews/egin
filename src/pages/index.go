package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully tested api",
	})
}
