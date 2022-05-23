package repository

import "github.com/go-redis/redis/v8"

func CreateRedisClient() *redis.Client {
	url := "redis://:@localhost:6262/0"
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
