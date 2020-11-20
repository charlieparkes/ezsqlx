package ezsqlx

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type ConnectionSettings struct {
	Host         string
	Port         string
	User         string
	Password     string
	Database     string
	DisableSSL   bool
	MaxOpenConns int
}

func (settings *ConnectionSettings) String() string {
	sslmode := "require"
	if settings.DisableSSL {
		sslmode = "disable"
	}
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Password,
		settings.Database,
		sslmode,
	)
}

func (settings *ConnectionSettings) Copy() *ConnectionSettings {
	cs := *settings
	return &cs
}

func (settings *ConnectionSettings) Connect() (*sqlx.DB, error) {
	return sqlx.Connect("pgx", settings.String())
}

func (settings *ConnectionSettings) Open() (*sqlx.DB, error) {
	return sqlx.Open("pgx", settings.String())
}

func (settings *ConnectionSettings) Init() *sqlx.DB {
	db, err := settings.Connect()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not connect to or ping database '%v': %v", settings.Database, settings.String()))
	}

	// Setting these arbitrarily to 50. We need to set a MaxOpenConns here since Oatmeal
	// uses channels during OrderProcessing and we don't want to hit any connection limits.
	db.SetMaxOpenConns(settings.MaxOpenConns)

	return db
}

func (settings *ConnectionSettings) Ping() error {
	var err error
	db, err := settings.Open()
	defer db.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Could not connect to %v: %v", settings.Host, err))
	}
	return db.Ping()
}

// Initialize databases given ConnectionStrings
func Init(databases map[string]*ConnectionSettings) (map[string]*sqlx.DB, func() error) {
	connections := map[string]*sqlx.DB{}

	for name, settings := range databases {
		connections[name] = settings.Init()
	}

	return connections, func() error {
		return CloseAll(connections)
	}
}

func CloseAll(connections map[string]*sqlx.DB) error {
	var err error
	for name, db := range connections {
		err = db.Close()
		if err != nil {
			log.WithFields(log.Fields{
				"event": "FatalDBCloseError",
				"error": err,
			}).Error(fmt.Sprintf("Could not close %v database", name))
		}
	}
	return err
}
