# NUT Forwarder for InfluxDB

<a href="https://hub.docker.com/r/albinodrought/nut-forwarder-influxdb">
  <img alt="albinodrought/nut-forwarder-influxdb Docker Pulls" src="https://img.shields.io/docker/pulls/albinodrought/nut-forwarder-influxdb">
</a>
<a href="https://github.com/AlbinoDrought/nut-forwarder-influxdb/blob/master/LICENSE">
  <img alt="AGPL-3.0 License" src="https://img.shields.io/github/license/AlbinoDrought/nut-forwarder-influxdb">
</a>

Forward some of your [Network UPS Tools (NUT)](https://networkupstools.org/index.html) data to InfluxDB. I built this for home use with my [CyberPower `CP1500AVRLCD`](https://networkupstools.org/ddl/Cyber_Power_Systems/CP1500AVRLCD.html), YMMV

## Running

```
INFLUX_SERVER=http://localhost:8086 \
NUT_HOST=localhost \
./nut-forwarder-influxdb
```

- `INFLUX_SERVER`: URL to InfluxDB server including scheme and port, defaults to `http://localhost:8086`

- `INFLUX_DATABASE`: Database to save data to, defaults to `ups`

- `INFLUX_USERNAME`: InfluxDB username, defaults to empty (no auth)

- `INFLUX_PASSWORD`: InfluxDB password, defaults to empty (no auth)

- `NUT_HOST`: hostname or IP address of your NUT server, defaults to `localhost`

- `NUT_USERNAME`: NUT username, defaults to empty (no auth)

- `NUT_PASSWORD`: NUT password, defaults to empty (no auth)


## Building

### Without Docker

```
go get -d -v
go build
```

### With Docker

`docker build -t albinodrought/nut-forwarder-influxdb .`

