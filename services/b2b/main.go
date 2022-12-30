package main

import (
	"b2b/containers"
	_ "go.uber.org/automaxprocs"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	if err := os.Setenv("TZ", "Asia/Bangkok"); err != nil {
		panic(err.Error())
	}
	app, err := containers.NewAppContainer()
	if err != nil {
		log.Fatalln(err)
	}
	_ = app.Start()
}
