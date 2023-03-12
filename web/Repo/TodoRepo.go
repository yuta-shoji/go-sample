package Repo

import (
	"fmt"
	"web/Http"
)

func TodoRepo() string {
	url := "http://localhost:8081/todos"
	channel := make(chan string)
	fmt.Println("request start")
	go Http.GetRequest(url, channel)
	fmt.Println("request end")

	fmt.Println(<-channel)
	return <-channel
}
