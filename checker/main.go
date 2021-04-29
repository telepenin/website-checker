package main

import (
	chckr "github.com/telepenin/website-checker/checker/checker"
	"github.com/telepenin/website-checker/checker/listener"
	config "github.com/telepenin/website-checker/config/src"
	"log"
)

func main() {

	cfg, err := config.FromEnv("CONFIG")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("init kafka listener")
	kafkaListener, err := listener.Init(cfg.Kafka)
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaListener.Close()

	checker := &chckr.Checker{
		Config: cfg.Checker,
		Processors: chckr.Processors{
			kafkaListener.Process,
		},
	}

	log.Println("run the website checker")
	checker.Run()

}
