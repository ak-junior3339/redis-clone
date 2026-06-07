package main

import (
	"fmt"
)

func main() {
	fmt.Println("Redis-Clone Starting...")
	store := NewStore()
	startServer(store)
}
