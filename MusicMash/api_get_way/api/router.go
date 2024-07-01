package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func RouterAPi() *gin.Engine {
	router := gin.Default()
	router.Group("api/")

	

}
