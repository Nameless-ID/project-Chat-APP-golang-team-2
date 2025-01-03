package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v8"
	"golang.org/x/net/context"
)

func CacheMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		key := c.Request.URL.Path

		// Cek cache
		val, err := redisClient.Get(ctx, key).Result()
		if err == nil {
			c.Data(http.StatusOK, "application/json", []byte(val))
			c.Abort()
			return
		}

		// Lanjutkan request
		c.Next()

		// Simpan ke cache
		status := c.Writer.Status()
		if status == http.StatusOK {
			body := c.Writer.Body
			redisClient.Set(ctx, key, body, 0) // Simpan tanpa TTL
		}
	}
}
