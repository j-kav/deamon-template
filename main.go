package main

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Do some async work.
	go p.run()
	return nil
}
func (p *program) run() {
	fmt.Println("run")
}
func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "{ServiceName}",
		DisplayName: "{ServiceName}",
		Description: "{ServiceName} service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}