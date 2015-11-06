package main

import (
	".."
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	// set configfile
	configfile := os.Getenv("SENSOR_CONFIG")
	if configfile == "" {
		// exist /etc/sensorserver/config.toml?
		if _, err := os.Stat("/etc/sensorserver/config.toml"); err == nil {
			configfile = "/etc/sensorserver/config.toml"
		}

		// only for testing: exist config.toml in the actual directory?
		if _, err := os.Stat("config.toml"); err == nil {
			configfile = "config.toml"
		}
	}

	// set data dir
	datadir := os.Getenv("SENSORSERVER_DATA")
	if datadir != "" {
		datadir += "/"
	}

	// get template dir
	templatedir := "/etc/sensorserver/templates"
	if _, err := os.Stat(templatedir); err != nil {
		templatedir = "templates"
	}

	log.Println("Read configuration form", configfile)
	s, port, err := sensorserver.New(configfile, datadir, templatedir)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.StaticFile("/favicon.ico", datadir+"resources/favicon.ico")
	router.StaticFile("/favicon-16x16.png", datadir+"resources/favicon-16x16.png")
	router.StaticFile("/favicon-32x32.png", datadir+"resources/favicon-32x32.png")
	router.StaticFile("/favicon-96x96.png", datadir+"resources/favicon-96x96.png")
	router.StaticFile("/android-icon-192x192.png", datadir+"resources/android-icon-192x192.png")

	router.LoadHTMLGlob(templatedir + "/*")
	router.Static("/assets", datadir+"assets")
	router.GET("/", s.Page)
	router.HEAD("/", s.Page)

	router.GET("/boltdb/backup", s.Backup)
	router.HEAD("/boltdb/backup", s.Backup)

	router.GET("/boltdb/stats", s.Stats)
	router.HEAD("/boltdb/stats", s.Stats)

	router.GET("/sensor", s.GetSensors)
	router.HEAD("/sensor", s.GetSensors)
	router.PUT("/sensor", s.PutMultiSensorData)
	router.PUT("/sensor/:name", s.PutSensorData)
	router.GET("/sensor/:name", s.GetSensorData)
	router.HEAD("/sensor/:name", s.GetSensorData)

	router.GET("/js/:type/chart.js", s.GetChart)
	router.HEAD("/js/:type/chart.js", s.GetChart)

	connect := fmt.Sprintf(":%d", port)
	router.Run(connect)

}
