package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const APP_CONFIG_NAME = "./configs/application.yml"

func main() {
	log.Printf("Application config %s. Load started", APP_CONFIG_NAME)
	appConfig, err := LoadAppConfig(APP_CONFIG_NAME)
	if err != nil {
		log.Fatalf("Application config %v. Load failed", APP_CONFIG_NAME)
		return;
	}
	serverConfig := appConfig.Server
	fmt.Printf("%+v", serverConfig)
	server := &http.Server{
		Addr: serverConfig.host + ":" + toStr(serverConfig.port),
		Handler: nil,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if (err != nil && err != http.ErrServerClosed) {
			log.Fatalf("Can't start server. Reason: %s", err.Error())
		}
	}()
	log.Printf("Server started at port %v", serverConfig.port)

	<-done
	log.Print("Stopping server at port ", serverConfig.port)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown failed:", err)
	}
	log.Print("Server stopped...")
}

func toStr(value uint16) string {
	return strconv.FormatUint(uint64(value), 10)
}
