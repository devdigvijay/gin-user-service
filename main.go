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
	"github.com/gin-gonic/gin"
)

func main() {
	// gin Engin as default /
	var ginEngin *gin.Engine = gin.Default()

	// User Controller /
	var userController controllers.UserController
	userController.Initialize(ginEngin)

	serve := &http.Server{
		Addr:    ":8080",
		Handler: ginEngin,
	}

	go func() {
		log.Println("🚀 Server started on :8080")
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v\n", err)
		}
	}()

	// Graceful shutdown /
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("🛑 Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if error := serve.Shutdown(ctx); error != nil {
		log.Println("error while shuting down!")
	}
	log.Println("✅ Server exited gracefully")

}
