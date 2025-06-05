package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Paraschamoli/students_API/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup router
	router:=http.NewServeMux()

	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})
	//setup server
	server:=http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler:router,
	}
	fmt.Printf("server started %s",cfg.HTTPServer.Address)
	err:=server.ListenAndServe()
	if err !=nil{
		log.Fatal("failed to start server")
	}

	
}