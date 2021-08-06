package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func createReids() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Print("ping error", err.Error())
		return nil
	}
	fmt.Print(pong, err)
	return client
}

func IsExist(name string) bool {
	client := createReids()
	a, err := client.Exists(name).Result()
	if err != nil {
		fmt.Println("判断key存在失败")
	}
	if a == 1 {
		fmt.Println("key存在")
		return true
	}
	return false
}

func SetValue(name string, value interface{}, time time.Duration) bool {

	client := createReids()
	err := client.Set(name, value, time).Err()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func GetValue(name string) (string, error) {

	client := createReids()
	value, err := client.Get(name).Result()
	if err != nil {
		fmt.Println("获取key失败")
		return "", err
	}
	fmt.Println(value)
	return value, err
}

func DeleteValue(name string) bool {

	client := createReids()
	statusDel, err := client.Del("ming").Result()

	if err != nil {
		return false
	}

	if 1 == statusDel {
		fmt.Println("删除值成功")
		return true
	}
	return false
}
