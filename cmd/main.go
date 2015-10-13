package main

import (
	"github.com/gin-gonic/gin"
	".."
    "log"
)

func main() {

    s, err := sensorserver.New("my_4.db")
    if err != nil {
        log.Fatal(err)
    }

	router := gin.Default()

	router.GET("/boltdb/backup", s.Backup)
	router.HEAD("/boltdb/backup", s.Backup)

    router.GET("/boltdb/stats", s.Stats)
    router.HEAD("/boltdb/stats", s.Stats)

    router.PUT("/sensor/:name", s.PutSensorData)

	router.Run(":8080")

}
