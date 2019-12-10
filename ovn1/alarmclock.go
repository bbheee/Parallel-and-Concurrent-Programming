package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	time.Sleep(delay * time.Second)
	h, m, _ := time.Now().Clock()
	fmt.Printf("Klockan är %d.%d: %s\n", h, m, text)
	Remind(text, delay)
}

func main() {
	go Remind("Dags att äta", 3)
	go Remind("Dags att arbeta", 8)
	go Remind("Dags att sova", 24)
	select {}
}
