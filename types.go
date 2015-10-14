package sensorserver

type putdata struct {
	Timestamp int
	Value     float32
}

type Config struct {
	Database string `toml:"database"`
	Title    string `toml:"title"`
}
