package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index API
func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "API works")
}
