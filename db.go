package main

import (
	_ "github.com/lib/pq"
	"os"
)

type DbConn struct {
	hostname string
	port string
	name string
	username string
	password string
}

// getConn gets the database connection details from environment variables
// If the documented environment variables are not set the method return default values
func (dc *DbConn) GetConn() DbConn {
	// set dc hostname
	if dc.hostname = os.Getenv("DB_HOSTNAME"); dc.hostname == "" {
		dc.hostname = "192.168.2.75"
	}
	// set dn port
	if dc.port = os.Getenv("DB_PORT"); dc.port == "" {
		dc.port = "5432"
	}
	// set dc name
	if dc.name = os.Getenv("DB_NAME"); dc.name == "" {
		dc.name = "assets"
	}
	// set dc username
	if dc.username = os.Getenv("DB_USERNAME"); dc.username == "" {
		dc.username = "postgres"
	}
	// set dc password
	if dc.password = os.Getenv("DB_PASSWORD"); dc.password == "" {
		dc.password = "postgres"
	}
	return *dc
}