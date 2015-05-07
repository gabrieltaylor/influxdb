package data

import (
	"github.com/influxdb/influxdb"
	"github.com/influxdb/influxdb/influxql"
)

// PayloadWriter accepts a payload from client facing endpoints such as
// HTTP JSON API, Collectd, Graphite, OpenTSDB, etc.
type PointsWriter interface {
	Write(*WritePointsRequest) error
}

type Queryer interface {
	Query(query *influxql.Query, databaseName string, user *influxdb.User, chunkeSize int) (chan *influxdb.Result, error)
}

type Authenticator interface {
	Authenticate(string, string) (*influxdb.User, error)
}

func NewDataNode() *Node {
	return &Node{}
}

type Node struct {
	//ms meta.Service
}

func (n Node) Write(r *WritePointsRequest) error {
	// 1. Check DB/RP/etc.. exists
	// 2. create Writers from metastore info
	return nil
}

// TODO implement
func (n Node) Authenticate(string, string) (*influxdb.User, error) {
	return nil, nil
}

// TODO implement
func (n Node) Query(query *influxql.Query, databaseName string, user *influxdb.User, chunkeSize int) (chan *influxdb.Result, error) {
	return nil, nil
}

func (n Node) getWriters(r *WritePointsRequest) ([]PointsWriter, error) {
	return nil, nil
}
