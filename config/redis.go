package config

import (
    "context"
    "github.com/redis/go-redis/v9"
    "log"
    "os"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: os.Getenv("REDIS_ADDR"),
        Password: "",
        DB: 0,
    })

    _, err := RedisClient.Ping(Ctx).Result()
    if err != nil {
        log.Fatal("Failed to connect to Redis")
    }
}
