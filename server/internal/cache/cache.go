package cache

import (
	"context"
	"log/slog"
	"os"

	"github.com/Shivraj1712/ClockOps.git/internal/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// @Summary 		Redis Connection
// @Description 	Used to setup and check for the stable connection between the server and the redis database
// @Tags			Database
// @Accept 			*/*
// @Produce 		plain
func ConnectRedis() {
	opts, err := redis.ParseURL(config.Configuration.RedisUrl)
	if err != nil {
		opts = &redis.Options{
			Addr: config.Configuration.RedisUrl,
		}
	}
	RedisClient = redis.NewClient(opts)
	ctx := context.Background()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		slog.Error("Redis database not connect")
		os.Exit(3)
	}
}
