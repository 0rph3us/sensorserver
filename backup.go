package sensorserver

import (
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Backup this method returns the hole database file
func (s *Sensorserver) Backup(c *gin.Context) {

	// filename contains the date
	time := time.Now()
	contentDisposition := `attachment; filename="` + time.Format("2006-01-02_15:04") + `-sensorserver.db"`

	err := s.boltdb.View(func(tx *bolt.Tx) error {
		c.Header("Content-Disposition", contentDisposition)
		c.Header("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(c.Writer)

		return err
	})
	if err != nil {

		c.String(http.StatusInternalServerError, err.Error())
	}

}
