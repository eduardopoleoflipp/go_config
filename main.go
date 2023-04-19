package main

import (
	"fmt"
	"github.com/eduardopoleoflipp/go_config"
)

func main() {
	var config = Config{name: "Main", url: "url.com"}
	fmt.Println(config.name)
}
