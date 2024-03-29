package db

import "flag"

// Host is the host of the database
var Host string

// Port is the port of the database
var Port string

// User is the user of the database
var User string

// Pass is the password of the database
var Pass string

// DB is the database to use
var DB string

// Options is the options for the database
type Options struct {
	Host string
	Port string
	User string
	Pass string
	DB   string
}

// NewDBOptions creates the new database options
func NewDBOptions() Options {
	return Options{
		Host: Host,
		Port: Port,
		User: User,
		Pass: Pass,
		DB:   DB,
	}
}

// Address compose the address of the database
func (o Options) Address() string {
	return o.Host + ":" + o.Port
}

func init() {
	flag.StringVar(&Host, "host", "localhost", "DB host.")
	flag.StringVar(&Port, "port", "3306", "DB port.")
	flag.StringVar(&User, "user", "root", "DB user.")
	flag.StringVar(&Pass, "password", "secret", "DB password.")
	flag.StringVar(&DB, "db", "book_management", "The DB name to use.")
}
