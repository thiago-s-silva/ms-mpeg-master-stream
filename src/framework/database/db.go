package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/domain"
)

// Database represents the database connection and configuration.
type Database struct {
	Db          *gorm.DB
	Dsn         string
	DsnTest     string
	DbType      string
	DbTypeTest  string
	Debug       bool
	AutoMigrate bool
	Env         string
}

// NewDb creates a new instance of the Database struct.
func NewDb() *Database {
	return &Database{}
}

// NewDbTest creates a new instance of the Database struct for testing purposes.
func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.Debug = true
	dbInstance.AutoMigrate = true

	conn, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	return conn
}

// Connect establishes a connection to the database.
func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env == "test" {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	} else {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrate {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		d.Db.Model(&domain.Job{}).AddForeignKey("video_id", "videos(id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil
}
