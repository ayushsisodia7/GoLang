package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Go redis tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ping)

	//Set single key pair
	// err = client.Set(context.Background(), "name", "Ayush", 0).Err()
	// if err != nil {
	// 	fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
	// 	return
	//}

	//Get the single key pair
	// val, err := client.Get(context.Background(), "name").Result()
	// if err != nil {
	// 	fmt.Printf("Failed to get the value from the redis instance: %s", err.Error())
	// 	return
	// }
	// fmt.Println("Value retrieved from Redis instance: ", val)

	//Set more information about a particular object
	type Person struct {
		ID         string
		Name       string
		Age        int
		Occupation string
	}

	UniqueID := uuid.NewString()

	JsonString, err := json.Marshal(Person{
		ID:         UniqueID,
		Name:       "Ayush",
		Age:        23,
		Occupation: "ASE",
	})

	if err != nil {
		fmt.Printf("failed to Marshal:%s", err.Error())
	}

	UniqueKey := fmt.Sprintf("person:%s", UniqueID)

	err = client.Set(context.Background(), UniqueKey, JsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set the value in the redis instance: %s", err.Error())
		return
	}

	//Get the information about a particular object
	val, err := client.Get(context.Background(), UniqueKey).Result()
	if err != nil {
		fmt.Printf("Failed to get the value from the redis instance: %s", err.Error())
		return
	}
	fmt.Println("Value retrieved from Redis instance: ", val)

}
