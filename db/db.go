package db

import (
	"fmt"
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

type Writer interface {
	Write(Point) error
}

type Queryer interface {
	Query(string) ([]influxdb.Result, error)
}

type Point interface {
	Name() string
	Tags() map[string]string
	Fields() map[string]interface{}
	Time() time.Time
}

// Create Database
// By default CREATE DATABASE uses IF NOT EXISTS
func Create(config Configurer) error {
	client, err := New(config)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Query(fmt.Sprintf("CREATE DATABASE \"%s\"", config.DB()))
	return err
}

type DB struct {
	// Exported
	// Unexported
	config Configurer
	client influxdb.Client
}

// Taken from https://github.com/influxdata/influxdb/tree/v0.13.0/client#querying-data
// with minor modifications
// Allows raw queries to database
func (db *DB) Query(cmd string) (res []influxdb.Result, err error) {
	log.WithField("query", cmd).Debug("query db")
	// Construct an InfluxDB Query
	qry := influxdb.Query{
		Command:  cmd,
		Database: db.config.DB(),
	}
	// Perform query on client
	rsp, err := db.client.Query(qry)
	if err != nil {
		return nil, err
	}
	// Error check the response
	if rsp.Error() != nil {
		return nil, rsp.Error()
	}
	return rsp.Results, nil
}

// Writes a new point to the DB
func (d *DB) Write(p Point) error {
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database: d.config.DB(),
	})
	if err != nil {
		return err
	}
	pt, err := influxdb.NewPoint(p.Name(), p.Tags(), p.Fields(), p.Time())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	return d.client.Write(bp)
}

// Close client
func (d *DB) Close() error {
	return d.client.Close()
}

// Constructs a new DB
func New(c Configurer) (*DB, error) {
	client, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr:     c.Address(),
		Username: c.Username(),
		Password: c.Password(),
	})
	if err != nil {
		return nil, err
	}
	return &DB{
		config: c,
		client: client,
	}, nil
}

// Simple with db decorator style, run any function with a fresh db connection
func With(fn func(db *DB) error) error {
	db, err := New(NewConfig())
	if err != nil {
		return err
	}
	defer db.Close()
	return fn(db)
}
