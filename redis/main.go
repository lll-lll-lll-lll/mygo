package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {

}

func connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		// DB:       0,  // use default DB
	})
	return rdb
}

func helloworld() {
	ctx := context.Background()
	rdb := connect()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
