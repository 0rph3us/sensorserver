package main

import (
	".."
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	// set configfile
	configfile := os.Getenv("SENSOR_CONFIG")
	if configfile == "" {
		configfile = "config.toml"
	}

	s, err := sensorserver.New(configfile)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.GET("/", s.Page)

	router.GET("/boltdb/backup", s.Backup)
	router.HEAD("/boltdb/backup", s.Backup)

	router.GET("/boltdb/stats", s.Stats)
	router.HEAD("/boltdb/stats", s.Stats)

	router.GET("/sensor", s.GetSensors)
	router.PUT("/sensor/:name", s.PutSensorData)
	router.GET("/sensor/:name", s.GetSensorData)

	router.Run(":8080")

}
