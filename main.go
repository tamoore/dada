package main

import (
	"fmt"
	"time"

	"github.com/tamoore/dada/internal/config"
)

func main() {
	c := make(chan config.Product)

	go config.Start(c)

	conf := <-c

	fmt.Println("Package manager: " + conf.PackageManager)
	fmt.Printf("Programs to install %v\n", conf.Dependencies)
	time.Sleep(time.Millisecond * 900)
}
