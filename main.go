package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-apiserver/config"
	"go-apiserver/router"
	"log"
	"net/http"
	"time"
)

var (
	addr = ":8080"
	cfg = pflag.StringP("config", "c", "",  "Apiserver config file path.")
)

func main() {
	// 解析命令行参数
	pflag.Parse()

	// init config.
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("addr"))
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	addr = viper.GetString("addr")
	log.Print("run mode:", viper.GetString("runmode"))

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

	log.Printf("Start to listening the incoming requests on http address: %s", addr)
	log.Printf(http.ListenAndServe(addr, g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error  {
	for i := 0; i < 3; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Println("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
