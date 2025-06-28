package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Paraschamoli/students_API/internal/config"
	"github.com/Paraschamoli/students_API/internal/http/handlers/student"
	"github.com/Paraschamoli/students_API/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	storage,err:=sqlite.New(cfg)
	if err!=nil{
	log.Fatal(err)
}
slog.Info("database initialized successfully")
	//setup router
	router:=http.NewServeMux()

	router.HandleFunc("POST /api/students",student.New(storage))
	router.HandleFunc("GET /api/students/{id}",student.GetById(storage))
	//setup server
	server:=&http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler:router,
	}

	done:=make(chan os.Signal,1)
	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
	go func() {
	fmt.Printf("server started %s\n", cfg.HTTPServer.Address)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}()

     <-done
	 slog.Info("shutting down the server")

	 ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	 defer cancel()

	 err=server.Shutdown(ctx)
	 if err!=nil{
		slog.Error("server shutdown failed", slog.String("error",err.Error()))
	} else {
		slog.Info("server gracefully stopped")
	}
	
}