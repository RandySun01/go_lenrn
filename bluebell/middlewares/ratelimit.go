package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

/*
@author RandySun
@create 2022-01-23-17:16
*/
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		//if bucket.Take(1) > 0{
		//
		//}
		if bucket.TakeAvailable(1) == 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}

		// 取到令牌就放行
		c.Next()
	}
}
