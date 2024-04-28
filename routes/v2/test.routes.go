package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TestRoutes(api *gin.RouterGroup) {
	api.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

}
