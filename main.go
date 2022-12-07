package main

import (
	"log"
	"ordersmod/api"
	"ordersmod/app"
	"time"

	"github.com/arshabbir/utils/config"
	"github.com/arshabbir/utils/logger"
	"github.com/arshabbir/utils/model"
)

func main() {
	log.Println("Starting the orders service....")

	conf := readConfig()
	l := logger.NewLogger(conf)
	app, err := app.NewApp(api.NewOrderServer(conf, l), conf, l)
	if err != nil {
		l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "main", Message: "Error initiliazing the app"})
		return
	}
	l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "main", Message: "App  initialization successful "})
	if err := app.StartApp(); err != nil {
		l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "main", Message: "Error starting the app"})
	}

}

func readConfig() *config.Config {
	return &config.Config{AppPort: 8080, LogLevel: 0}
}
