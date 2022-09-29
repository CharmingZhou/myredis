package main

import (
	"fmt"

	"github.com/CharmingZhou/myredis/tcp"
)

var banner = `
	myredis
`

func main() {
	fmt.Printf("%s start\n", banner)
	err := tcp.ListenAndServeWithSignal(&tcp.Config{
		Address:    "127.0.0.1:6379",
		MaxConnect: 100,
		Timeout:    1000,
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
}
