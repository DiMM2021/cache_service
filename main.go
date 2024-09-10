package main

import (
	"cache_service/cache"
	"fmt"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")
	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")
	fmt.Println(userId)
}
