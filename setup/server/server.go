package server

import (
	"context"
	"flight/setup/database"
	"flight/setup/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

const TIMEOUT = 5

func Run() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connect DB: %s", err)
	}

	router := router.NewRouter(db)

	s := &http.Server{
		Addr:         os.Getenv("API_PORT"),
		Handler:      router,
		ReadTimeout:  TIMEOUT * time.Second,
		WriteTimeout: TIMEOUT * time.Second,
	}

	go func() {
		if err = s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}

	<-ctx.Done()

	log.Println("Server exits successfully")
}
