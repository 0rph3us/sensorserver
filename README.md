# Sensorserver

Sensorserver is simple all-in-one solution for store and
serve timeseries from sensors. I develop this application
for my Raspberry Pi.

## Build

``` sh
./prepare.sh
go get -u github.com/0rph3us/astrotime
go get -u github.com/BurntSushi/toml
go get -u github.com/gin-gonic/gin
go get -u github.com/boltdb/bolt

go run cmd/main.go
```

## Dependencies
* [Gin] the web framework
* [toml] is the format for the configuration file
* [BoltDB] the storage backend
* [Astrotime] for calculate the sunset and sunrise

## API

### /boltdb/backup
Support GET and HEAD requests. Get a binary stream of the database file.
``` sh
curl -J -O http://127.0.0.1:8080/boltdb/backup
```

### /boltdb/stats
Support GET and HEAD requests. The response is a JSON.
``` sh
curl 127.0.0.1:8080/boltdb/stats?pretty=true
```

### /sensor
Support GET, HEAD and PUT requests. The GET response is
JSON Array with the name of all sensors. The PUT method put
date from different sensors into to storage backend.
``` sh
curl -XPUT 'http://127.0.0.1:8080/sensor' -d '[
    {
        "Timestamp": 1445311973,
        "Sensors": {
            "temp_bmp085": 23.5,
            "humidity": 43.7,
            "pressure": 1004.53467
        }
    },
    {
        "Timestamp": 1445312973,
        "Sensors": {
            "temp_bmp085": 23.4,
            "humidity": 43.6,
            "pressure": 1002.8329
        }
    }
]'
```

### /sensor/:name

Support GET, HEAD and PUT requests for a specific sensor.
GET returns a JSON array with timestamp-value objects. The
timestamp is in seconds since 01.01.1970
``` sh
curl -XGET 'http://127.0.0.1:8080/sensor/pressure?pretty=true'
[
    {
        "Timestamp": 1413011716487,
        "Value": 1003.35
    },
    {
        "Timestamp": 1413076519461,
        "Value": 1003.17896
    },
    {
        "Timestamp": 1413141322435,
        "Value": 1002.38873
    },
    {
        "Timestamp": 1413206125409,
        "Value": 996.93915
    }
]
```

The PUT method put new data into the server. The timestamp
is a normal 32 Bit Unixtimestamp.
``` sh
curl -XPUT 'http://127.0.0.1:8080/sensor/humidity' -d '[
    {
        "Timestamp": 1412979315,
        "Value": 57.9
    },
    {
        "Timestamp": 1412979343,
        "Value": 58
    },
    {
        "Timestamp": 1412979394,
        "Value": 58
    },
    {
        "Timestamp": 1412979407,
        "Value": 58.1
    },
    {
        "Timestamp": 1412979417,
        "Value": 57.9
    },
    {
        "Timestamp": 1412979437,
        "Value": 58.1
    },
    {
        "Timestamp": 1412979557,
        "Value": 58.3
    }
]'
```

## Alternative

* [influxDB]
* [prometheus]
* [graphite]

All alternatives looks good. Graphite has the most functions, but the project looks
dead. [Grafana] is a very good frontend for the 3 backends. [This plugin] is required
for the prometheus datasource.

[Gin]: https://github.com/gin-gonic/gin/
[toml]: https://github.com/BurntSushi/toml/
[BoltDB]: https://github.com/boltdb/bolt/
[Grafana]: http://grafana.org/
[influxDB]: https://influxdb.com/
[graphite]: http://graphite.wikidot.com/
[Astrotime]: https://github.com/0rph3us/astrotime/
[prometheus]: http://prometheus.io/
[This plugin]: https://github.com/grafana/grafana-plugins/tree/master/datasources/prometheus
