package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func isPrime(value int) bool {
	for i := 2; i <= value/2; i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(
		render.Options{
			Directory: "templates",
		},
	))

	m.Get("/500", func(r render.Render, req *http.Request) {
		r.Error(500)
	})
	m.Get("/", func(r render.Render, req *http.Request) {
		log.Printf("ENVIRONMENT %+v\n", os.Environ())
		wait := req.URL.Query().Get("wait")
		if wait != "" {
			sleep, err := time.ParseDuration(wait)
			if err != nil {
				r.Error(400)
				return
			}
			log.Printf("Sleep for %s seconds\n", wait)
			time.Sleep(sleep)
		}
		if req.URL.Query().Get("prime") != "" {
			val, _ := strconv.Atoi(req.URL.Query().Get("prime"))
			log.Printf("Is %d prime: %t", val, isPrime(val))
		}
		r.HTML(200, "index", nil)
	})

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	go http.Serve(listener, m)
	log.Println("Listening on 0.0.0.0:" + port)

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
	listener.Close()
}
