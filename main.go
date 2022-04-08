package main

import (
	"fmt"
	"github.com/kardianos/service"
	"time"
)

const serviceName = "My Service"
const serviceDescription = "It's my learning course!"

type program struct {
}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " Started")
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	fmt.Println(s.String() + " Stopped")
	return nil
}

func (p program) run() {
	for {
		fmt.Println("Service is running")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}

	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service", err.Error())
	}

	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service", err.Error())
	}
}
