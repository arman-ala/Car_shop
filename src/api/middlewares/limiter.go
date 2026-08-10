package middlewares

import (
	"net/http"

	tlb "github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequestMiddleware() gin.HandlerFunc {
	lmt := tlb.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tlb.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
