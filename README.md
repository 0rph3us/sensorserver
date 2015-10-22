# Sensorserver

Sensorserver is simple all-in-one solution for store and
serve timeseries from sensors. I develop this application
for my Raspberry Pi.

## Dependencies
* [Gin] the web framework
* [toml] is the format for the configuration file
* [BoltDB] the storage backend
* [Astrotime] for calculate the sunset and sunrise

## API

### /boltdb/backup
Support GET and HEAD requests. Get a binary stream of the database file.
``` sh
curl 127.0.0.1:8080/boltdb/backup > backup$(date +%F).db
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
timestamp is in milliseconds since 01.01.1970 (needed for
javascript)
``` sh
curl -XGET 'http://127.0.0.1:8080/sensor/pressure'
[
    {
        "T": 1413011716487,
        "V": 1003.35
    },
    {
        "T": 1413076519461,
        "V": 1003.17896
    },
    {
        "T": 1413141322435,
        "V": 1002.38873
    },
    {
        "T": 1413206125409,
        "V": 996.93915
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

[Gin]: https://github.com/gin-gonic/gin/
[toml]: https://github.com/BurntSushi/toml/
[BoltDB]: https://github.com/boltdb/bolt/
[Astrotime]: https://github.com/0rph3us/astrotime/
