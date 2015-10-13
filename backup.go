package sensorserver

import (
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Sensorserver) Backup(c *gin.Context) {

	err := s.boltdb.View(func(tx *bolt.Tx) error {
		c.Header("Content-Disposition", `attachment; filename="my.db"`)
		c.Header("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(c.Writer)

		return err
	})
	if err != nil {

		c.String(http.StatusInternalServerError, err.Error())
	}

}
