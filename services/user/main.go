package main

import (
	_ "go.uber.org/automaxprocs"
	"log"
	"math/rand"
	"os"
	"time"
	"user/containers"
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