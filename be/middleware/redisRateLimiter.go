package middleware

import (
	"net/http"
	"strconv"
	"time"
	"zydukas-rankings/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func RedisConfig() (*redis.Options, error) {
	redisAddress, err := utils.EnvString("REDIS_ADDRESS")
	if err != nil {
		return nil, err
	}

	return &redis.Options{Addr: redisAddress}, nil
}

func RedisMiddeware(rdb *redis.Client, window time.Duration) gin.HandlerFunc {
	limit := int64(10)
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		key := "ratelimit:" + c.ClientIP()

		pipe := rdb.TxPipeline()
		count := pipe.Incr(ctx, key)
		pipe.ExpireNX(ctx, key, window) // start the window on the first request only
		if _, err := pipe.Exec(ctx); err != nil {
			zap.S().Warnw("rate limit check failed", "error", err)
			c.Next() // if Redis is down, let requests through instead of blocking everyone
			return
		}

		if count.Val() > limit {
			ttl, _ := rdb.TTL(ctx, key).Result()
			c.Header("Retry-After", strconv.Itoa(int(ttl.Seconds())))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}

		c.Next()
	}
}
