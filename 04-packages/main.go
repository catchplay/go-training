package main

import (
	"fmt"
	"time"

	"github.com/catchplay/go-training/04-packages/user"
)

func main() {
	now := time.Now()
	fmt.Printf("current time: %s\n", now.Format(time.RFC3339))
	fmt.Printf("%s", user.SayHello())
}
