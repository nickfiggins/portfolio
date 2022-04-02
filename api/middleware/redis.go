package middleware

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var ctx = context.Background()
var Redis *redis.Client

func RunRedis() {
    rdb := redis.NewClient(&redis.Options{
        Network:  "tcp",
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
        MaxRetries: 3,
    })
	fmt.Println(rdb.Ping().Result())
    err := rdb.Set("key", "value", 0).Err()
    if err != nil {
        panic(err)
    }
	Redis = rdb
}