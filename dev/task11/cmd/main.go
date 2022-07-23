package main

import (
	"context"
	calendar "github.com/infamax/l2/task11"
	"github.com/infamax/l2/task11/config"
	"github.com/infamax/l2/task11/internal/handlers"
	"github.com/infamax/l2/task11/internal/repository"
	"github.com/infamax/l2/task11/internal/service"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	fileBytes, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("cannot read file")
	}

	cf, err := config.ParseConfig(fileBytes)
	if err != nil {
		log.Fatal("cannot parse config")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	db, err := repository.New(ctx, cf.GetConnectionString())
	if err != nil {
		log.Fatal("cannot connect to d")
	}

	s, err := service.New(db)

	if err != nil {
		log.Fatal("cannot create s")
	}

	handler, err := handlers.New(s)

	if err != nil {
		log.Fatal("cannot create handler")
	}

	srv := new(calendar.Server)
	go func() {
		if err := srv.Run(strconv.Itoa(cf.Port), handler.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("CalendarApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("CalendarApp finished")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

}
