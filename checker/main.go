package main

import (
	chckr "github.com/telepenin/website-checker/checker/checker"
	"github.com/telepenin/website-checker/checker/listener"
	"github.com/telepenin/website-checker/shared"
	"log"
)

func main() {

	cfg, err := shared.ConfigFromEnv("CONFIG")
	if err != nil {
		log.Fatal(err)
	}

	stdoutListener := &listener.StdoutListener{}

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
			stdoutListener.Process,
		},
	}

	log.Println("run the website checker")
	checker.Run()

}
