package main

import (
	"errors"
	"go-apiserver/router"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

var port = ":8080"


func main()  {
	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}
	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlewares.
		middlewares...,
		)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Println("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", port)
	log.Printf(http.ListenAndServe(port, g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error  {
	for i := 0; i < 3; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Println("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
