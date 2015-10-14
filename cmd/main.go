package main

import (
	".."
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	s, err := sensorserver.New("my_4.db")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Static("/assets", "./assets")

	router.GET("/boltdb/backup", s.Backup)
	router.HEAD("/boltdb/backup", s.Backup)

	router.GET("/boltdb/stats", s.Stats)
	router.HEAD("/boltdb/stats", s.Stats)

	router.GET("/sensor", s.GetSensors)
	router.PUT("/sensor/:name", s.PutSensorData)
	router.GET("/sensor/:name", s.GetSensorData)

	router.Run(":8080")

}
