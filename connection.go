package ezsqlx

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type ConnectionSettings struct {
	Driver       string
	Host         string
	Port         string
	User         string
	Password     string
	Database     string
	DisableSSL   bool
	MaxOpenConns int
}

func (cs *ConnectionSettings) String() string {
	sslmode := "require"
	if cs.DisableSSL {
		sslmode = "disable"
	}
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cs.Host,
		cs.Port,
		cs.User,
		cs.Password,
		cs.Database,
		sslmode,
	)
}

func (cs *ConnectionSettings) Copy() *ConnectionSettings {
	s := *cs
	return &s
}

func (cs *ConnectionSettings) getDriver() string {
	if cs.Driver == "" {
		return "pgx"
	}
	return cs.Driver
}

func (cs *ConnectionSettings) Connect() (*sqlx.DB, error) {
	return sqlx.Connect(cs.getDriver(), cs.String())
}

func (cs *ConnectionSettings) Open() (*sqlx.DB, error) {
	return sqlx.Open(cs.getDriver(), cs.String())
}

func (cs *ConnectionSettings) Init() *sqlx.DB {
	db, err := cs.Connect()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not connect to or ping database '%v': %v", cs.Database, cs.String()))
	}

	if cs.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cs.MaxOpenConns)
	}

	return db
}

func (cs *ConnectionSettings) Ping() error {
	var err error
	db, err := cs.Open()
	if err != nil {
		return fmt.Errorf("could not connect to %v: %v", cs.Host, err)
	}
	defer db.Close()
	return db.Ping()
}
