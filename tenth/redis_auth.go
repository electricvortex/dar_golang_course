package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"time"
)

const (
	fullControl = iota
	read
	write
	update
)

type Admin struct {
	permission int
}

type userInfo struct {
	Token string
	Admin Admin
}

func main() {
	userID := "10"
	redisClient := Connect()
	token := uuid.New()
	fmt.Println(token)
	input := &userInfo{Admin: Admin{permission:fullControl}, Token: token.String()}
	fmt.Println(input)
	data, err := json.Marshal(input)
	err = redisClient.Set("token: " + token.String(), data, 1 * time.Minute).Err()
	if err != nil {
		panic(err)
	}

	//val, err := redisClient.Get("token: " + userID).Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("token for user: " + userID + " - ", val)

	val2, err := redisClient.Get("token: " + token.String()).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		uInfo := &userInfo{}
		err := json.Unmarshal([]byte(val2), uInfo)
		if err != nil {
			panic(err)
		}
		fmt.Println(uInfo)
		fmt.Printf("token for user " + userID + "is: %v", uInfo.Token)
	}

}

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

