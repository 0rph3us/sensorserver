package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

import _ "github.com/go-sql-driver/mysql"

type data struct {
	Timestamp int
	Value     float32
}

func writeFile(d []data, name string) {
	var j []byte

	j, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := ioutil.WriteFile(name, j, 0644); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var insert_time int
	var temp_bmp085, tmp_dth22, pressure, p_sea, humidity float32

	u, _ := user.Current()

	user := flag.String("user", u.Username, "user for connect to database")
	pass := flag.String("password", "", "password for the user")
	prot := "tcp"
	host := flag.String("host", "localhost", "hostname")
	port := flag.Int("port", 3306, "MySQL port")
	dbname := flag.String("database", "sensoren", "name of the database")

	flag.Parse()

	netAddr := fmt.Sprintf("%s(%s:%d)", prot, *host, *port)
	dsn := fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true&parseTime=true", *user, *pass, netAddr, *dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	stmt, err := db.Prepare("SELECT UNIX_TIMESTAMP(insert_time), temp_bmp085, tmp_dth22, pressure, p_sea, humidity FROM test WHERE insert_time > ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query("2012-10-10 22:30:05")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var j_temp_bmp085, j_tmp_dth22, j_pressure, j_p_sea, j_humidity []data
	for rows.Next() {
		err := rows.Scan(&insert_time, &temp_bmp085, &tmp_dth22, &pressure, &p_sea, &humidity)
		if err != nil {
			log.Fatal(err)
		}

		j_temp_bmp085 = append(j_temp_bmp085, data{insert_time, temp_bmp085})
		j_tmp_dth22 = append(j_tmp_dth22, data{insert_time, tmp_dth22})
		j_pressure = append(j_pressure, data{insert_time, pressure})
		j_p_sea = append(j_p_sea, data{insert_time, p_sea})
		j_humidity = append(j_humidity, data{insert_time, humidity})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	writeFile(j_temp_bmp085, "temp_bmp085")
	writeFile(j_tmp_dth22, "tmp_dth22")
	writeFile(j_pressure, "pressure")
	writeFile(j_p_sea, "p_sea")
	writeFile(j_humidity, "humidity")
}
