package main

import (
	"MobeReviewAPI/Internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
