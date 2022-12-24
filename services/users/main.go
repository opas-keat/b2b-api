package main

import (
	"log"
	"math/rand"
	"time"
	"users/containers"

	_ "go.uber.org/automaxprocs"
)

func main() {
	rand.Seed(time.Now().Unix())
	app, err := containers.NewAppContainer()
	if err != nil {
		log.Fatalln(err)
	}
	_ = app.Start()
}
