package main

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"github.com/wzslr321/database/postgres"
	redisdb "github.com/wzslr321/database/redis"
	"github.com/wzslr321/routers"
	"github.com/wzslr321/settings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	settings.InitSettings()
	postgres.InitPostgre()
	redisdb.InitRedis()
}

func main() {

	addr := fmt.Sprintf(":%s", settings.ServerSettings.Addr)
	port := addr
	router := cors.Default().Handler(routers.InitRouter())
	readTimeout := settings.ServerSettings.ReadTimeout
	writeTimeout := settings.ServerSettings.WriteTimeout
	mxHdrBytes := settings.ServerSettings.MaxHeaderBytes

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: mxHdrBytes,
	}

	var err error

	err = s.ListenAndServe()

	log.Printf("Server is running on port:  %s", port)

	go func() {
		if err = s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		// Waiting for interrupt by system signal
		<-quit
		time.Sleep(180 * time.Second)
		if err = s.Shutdown(ctx); err != nil {
			log.Fatal("Server shutdown", err)
		}

		select {
		case <-ctx.Done():
			log.Println("Timeout of 5 seconds")
		}
		log.Println("Server exiting")
		os.Exit(1)

	}()

	log.Println("Shutting down the server...")

}
