package src

import (
	"egin/src/pages"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/", pages.IndexPage)

	err := r.Run()

	if err != nil {
		return
	}
}
