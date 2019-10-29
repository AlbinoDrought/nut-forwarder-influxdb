package main

type config struct {
	InfluxServer   string `env:"INFLUX_SERVER" envDefault:"http://localhost:8086"`
	InfluxDatabase string `env:"INFLUX_DATABASE" envDefault:"ups"`
	InfluxUsername string `env:"INFLUX_USERNAME"`
	InfluxPassword string `env:"INFLUX_PASSWORD"`

	NUTHost     string `env:"NUT_HOST" envDefault:"localhost"`
	NUTUsername string `env:"NUT_USERNAME"`
	NUTPassword string `env:"NUT_PASSWORD"`
}
