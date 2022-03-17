package models

import (
	"fmt"
	"time"
	"strings"

	mysql "github.com/go-sql-driver/mysql"
	//"cloud.google.com/go/cloudsqlconn"
)

// DbConnConfig holds all info necessary to make a DSN connection string for a mysql VIP
type DbConnConfig struct {
	Driver			string
	Host 			string
	Port			int32
	User       		string
	Password   		string
	Timeout    		time.Duration
	DBName			string
	DSN        		string
	InstanceConnection	string
}

// MakeDSN makes a mysql DSN connection string in the DbConnConfig struct
func (d *DbConnConfig) MakeDSN() {
	dsnConfig := mysql.NewConfig()
	dsnConfig.User = d.User
	dsnConfig.Passwd = d.Password
	dsnConfig.Timeout = d.Timeout
	dsnConfig.Net = "tcp"
	dsnConfig.DBName = d.DBName
	dsnConfig.Addr = fmt.Sprintf("%s:%d", d.Host, d.Port)
	d.DSN = dsnConfig.FormatDSN()
}

func (d *DbConnConfig) MakeInstanceDSN() {
	dsnConfig := mysql.NewConfig()
	dsnConfig.User = d.User
	dsnConfig.Passwd = d.Password
	dsnConfig.Timeout = d.Timeout
	dsnConfig.Net = d.Driver
	dsnConfig.DBName = d.DBName
	dsnConfig.Addr = fmt.Sprintf("%s", d.InstanceConnection)
	d.DSN = dsnConfig.FormatDSN()
}

func (d *DbConnConfig) String() string {
	// Strip mysql passwords before dumping
	d.Password = "<PASSWORD>"
	return strings.ReplaceAll(d.DSN, d.Password, "<PASSWORD>")
}
