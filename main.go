package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kardianos/service"
	"net/http"
	"os"
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
	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))

	//router.GET("/", serveHomepage)

	err := http.ListenAndServe(":81", router)
	if err != nil {
		fmt.Println("Problem starting web server: " + err.Error())
		os.Exit(-1)
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
