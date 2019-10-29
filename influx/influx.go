package influx

import (
	"fmt"

	influx "github.com/influxdata/influxdb1-client/v2"
)

// Options specifies how to connect to our InfluxDB server
type Options struct {
	Server   string
	Database string
	Username string
	Password string
}

// Influxable things can be sent to InfluxDB
type Influxable interface {
	Tags() map[string]string
	Fields() map[string]interface{}
	Category() string
}

// Client automagically sends your Influxable resources
type Client struct {
	client  *influx.Client
	options *Options
}

// Connect to the target InfluxDB server
func Connect(options Options) (*Client, error) {
	config := influx.HTTPConfig{
		Addr: options.Server,
	}

	if options.Username != "" || config.Password != "" {
		config.Username = options.Username
		config.Password = options.Password
	}

	client, err := influx.NewHTTPClient(config)
	if err != nil {
		return nil, err
	}

	// attempt to create DB
	query := influx.NewQuery(fmt.Sprintf("CREATE DATABASE %v", options.Database), "", "")
	_, err = client.Query(query)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:  &client,
		options: &options,
	}, nil
}

// Send an Influxable thing to the configured InfluxDB
func (influxClient *Client) Send(thing Influxable) error {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  influxClient.options.Database,
		Precision: "s",
	})

	if err != nil {
		return err
	}

	pt, err := influx.NewPoint(thing.Category(), thing.Tags(), thing.Fields())

	if err != nil {
		return err
	}

	bp.AddPoint(pt)

	return (*influxClient.client).Write(bp)
}
