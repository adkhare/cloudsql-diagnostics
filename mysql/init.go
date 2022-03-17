package mysql

import (
	"cloud.google.com/go/cloudsqlconn"
	mysqlcloudsqlconn "cloud.google.com/go/cloudsqlconn/mysql/mysql"
)

type Mysql struct {
	Driver string
}

// Initiatize diagnostics for the said mysql driver eg. cloudsqlproxy
func (m *Mysql) Init(opts ...cloudsqlconn.Option) error {
	cleanup, err := mysqlcloudsqlconn.RegisterDriver(m.Driver,opts...)
    defer cleanup()

	return err
}
