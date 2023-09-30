package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"task-manager/app/infrastructure"
	"time"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	router := infrastructure.NewRouter()
	router.Start()

	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	router.Shutdown(ctx)
}
