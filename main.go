package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devdigvijay/gin-user-service/controllers"
	"github.com/devdigvijay/gin-user-service/environment"
	"github.com/devdigvijay/gin-user-service/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration, error := environment.Load(utils.LoadEnvFlags())
	if error != nil {
		log.Fatal("---> error while parsing env")
	}
	// gin Engin as default /
	var ginEngin *gin.Engine = gin.Default()

	// User Controller /
	var userController controllers.UserController
	userController.Initialize(ginEngin)

	serve := &http.Server{
		Addr:    configuration.App.Port,
		Handler: ginEngin,
	}

	go func() {
		log.Println("---> Server started on " + configuration.App.Port)
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("---> listen error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configuration.Server.ShutDownTimeout))
	defer cancel()

	if error := serve.Shutdown(ctx); error != nil {
		log.Println("---> error while shuting down!")
	}
	log.Println("---> Server exited gracefully")

}
