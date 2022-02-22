package xsingleflight

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

var sg singleflight.Group

func GetCache() (string, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 0,
	})
	key := "hello"
	ret := client.Get(ctx, key)
	return ret.Val(), nil
}

func SetCache() {
	fmt.Printf("调用了 SetCache\n")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 0,
	})
	key := "hello"
	client.Set(ctx, key, "WORLD", 10 * time.Second)
}

func GetData() (string, error) {
	val, err := GetCache()
	if err != nil {
		return "", err
	}
	if val != "" {
		return val, nil
	}
	va, er, share := sg.Do("hello", func() (interface{}, error) {
		SetCache()
		return "WORLD", nil
	})
	if er != nil {
		return "", err
	}
	fmt.Printf("share:%v\n", share)
	return va.(string), nil
}

func ManyThread()  {
	num := 10
	wait := sync.WaitGroup{}
	wait.Add(num)
	for i:=0; i<= num; i++ {
		go func() {
			defer wait.Done()
			val, err := GetData()
			fmt.Printf("val:%s err:%v\n", val, err)
		}()
	}
	wait.Wait()
}





