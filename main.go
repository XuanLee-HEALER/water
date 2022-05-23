package main

import (
	"context"
	"fmt"
	"log"

	"github.com/XuanLee-HEALER/water/core"
	"github.com/XuanLee-HEALER/water/repository"
)

type MyWork struct{}

func (w MyWork) Work() {
	fmt.Println("my work done!")
}

func main() {
	f := core.CreateWorkFlow(MyWork{})
	f.StartWork()

	ctx := context.Background()

	rdb := repository.CreateRedisClient()
	// rdb.Set(ctx, "name", "lixuan", 0)
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		log.Printf("key:%v not exists.\n", "name")
	}
	fmt.Println("name => ", val)
}
