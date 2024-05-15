package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Go Redis Basic:")

	//NewClient returns a client to the Redis Server specified by Options.
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//connect to container
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(ping)

	type Person struct {
		ID          string
		Name        string `json:"name"`
		Age         int    `json:"age"`
		Ocucupation string `json:"occupation"`
	}

	//Marshal returns the JSON encoding of v
	// {"ID":"a934a569-06dc-40a4-a189-9a674319deee","name":"Jorge Eliecer","age":33,"occupation":"Software Engineer"}

	jorgeID := uuid.NewString()
	jsonString, err := json.Marshal(Person{
		ID:          jorgeID,
		Name:        "Jorge Eliecer",
		Age:         33,
		Ocucupation: "Software Engineer",
	})
	if err != nil {
		fmt.Println("Error on Marshal: ", err.Error())
		return
	}

	jorgeKey := fmt.Sprintf("person:%s", jorgeID)

	// context, key, value, expiration
	err = client.Set(context.Background(), jorgeKey, jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Fail to set value in the redis instance %s: ", err.Error())
		return
	}

	value, err := client.Get(context.Background(), jorgeKey).Result()
	if err != nil {
		fmt.Printf("Fail to get value from the redis instance %s: ", err.Error())
		return
	}
	fmt.Printf("Value retrieved from redis: %s\n", value)
}
