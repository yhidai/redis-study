package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var (
	endpoint string
	read     bool
)

func main() {
	flag.StringVar(&endpoint, "endpoint", "", "set redis endpoint. e.g. 10.41.0.42:31362")
	flag.BoolVar(&read, "read", false, "read only flag")
	flag.Parse()

	rdb := redis.NewClient(
		&redis.Options{
			Addr: endpoint,
		},
	)

	if !read {
		err := rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}
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
	// Output: key value
	// key2 does not exist
}
