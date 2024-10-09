package main

import (
	advance "cache/advance"
	basic "cache/basic"
	moderate "cache/moderate"
	"fmt"
)

func main() {
	fmt.Println("Machine Coding Round - Go")

	// basic
	basic.BasicRedisCache()

	//Moderate
	moderate.ModerateLRUCache()

	// Advanced
	advance.AdvanceLFUCache()

}
