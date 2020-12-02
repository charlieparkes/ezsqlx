package test

import (
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	// "github.com/ory/dockertest/v3/docker"
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

func (settings *ConnectionSettings) Connect() (*sqlx.DB, error) {
	return sqlx.Connect("pgx", settings.String())
}

func (settings *ConnectionSettings) Open() (*sqlx.DB, error) {
	return sqlx.Open("pgx", settings.String())
}

type PostgresDatabase struct {
	Settings *ConnectionSettings
	Pool     *dockertest.Pool
	Network  *dockertest.Network
	Resource *dockertest.Resource
}

func (d *PostgresDatabase) SetUp() error {
	var err error
	name := "ezsqlx_" + namesgenerator.GetRandomName(0)
	d.Settings = &ConnectionSettings{
		Host:       "localhost",
		User:       "postgres",
		Password:   name,
		Database:   name,
		DisableSSL: true,
	}

	d.Pool, err = dockertest.NewPool("")
	if err != nil {
		return err
	}

	d.Network, err = d.Pool.CreateNetwork(name)
	if err != nil {
		return err
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=" + d.Settings.User,
			"POSTGRES_PASSWORD=" + d.Settings.Password,
			"POSTGRES_DB=" + d.Settings.Database,
		},
		Networks: []*dockertest.Network{
			d.Network,
		},
	}
	d.Resource, err = d.Pool.RunWithOptions(&opts)
	if err != nil {
		return err
	}

	// Tell docker to kill the container after an unreasonable amount of test time to prevent orphans.
	d.Resource.Expire(600)

	waitForPostgres(d.Pool, d.Resource, d.Settings)

	return nil
}

func (d *PostgresDatabase) TearDown() error {
	err := d.Network.Close()
	if err != nil {
		return err
	}
	err = d.Pool.Purge(d.Resource)
	return err
}

func waitForPostgres(pool *dockertest.Pool, resource *dockertest.Resource, settings *ConnectionSettings) {
	if err := pool.Retry(func() error {
		var err error
		settings.Port = resource.GetPort("5432/tcp")
		db, err := sqlx.Open("pgx", settings.String())
		if err != nil {
			fmt.Printf("%v\n", err)
			return err
		}
		defer db.Close()
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
}

type FooBar struct {
	Id      int        `db:"id"`
	Message string     `db:"message"`
	Flip    bool       `db:"flip"`
	Created *time.Time `db:"created"`
}

func CreateDummyTable(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec(`CREATE TABLE foobar (
		id SERIAL PRIMARY KEY,
		message TEXT,
		flip BOOL NOT NULL DEFAULT false,
		created TIMESTAMP DEFAULT NOW()
	)`)
	tx.Commit()
}
