package db

import (
	"fmt"
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

type DB struct {
	// Exported
	// Unexported
	config Configurer
	client influxdb.Client
}

// Create Database
// By default CREATE DATABASE uses IF NOT EXISTS
func (db *DB) Create() error {
	_, err := db.Query(fmt.Sprintf("CREATE DATABASE \"%s\"", db.config.DB()))
	return err
}

// Taken from https://github.com/influxdata/influxdb/tree/v0.13.0/client#querying-data
// with minor modifications
// Allows raw queries to database
func (db *DB) Query(cmd string) (res []influxdb.Result, err error) {
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

// Write a new score point
func (d *DB) InsertScore(user string, score int) error {
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database: d.config.DB(),
	})
	if err != nil {
		return err
	}
	tags := map[string]string{"user": user}
	fields := map[string]interface{}{
		"score": score,
	}
	pt, err := influxdb.NewPoint("score", tags, fields, time.Now().UTC())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	return d.client.Write(bp)
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
