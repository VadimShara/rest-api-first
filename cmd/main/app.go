package main

import (
	"github.com/VadimShara/rest-api-tutorial/pkg/logging"
	"github.com/VadimShara/rest-api-first/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main(){
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router){
	logger := logging.GetLogger

	logger.Info("start application")

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	logger.Info("server is listening part 0.0.0.0:1234")
	logger.Fatalln(server.Serve(listener))
}