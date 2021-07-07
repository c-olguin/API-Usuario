package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	// database configuration
	Username string
	Password string
	Hostname string
	Database string

	// database properties
	MaxIddleConn    int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

func getDataBaseSpecification() DatabaseConfig {
	var databaseConfig DatabaseConfig

	//if os.Getenv("ENVIRONMENT") == "production" {
		databaseConfig = DatabaseConfig{
			Username:        "root",
			Password:        "Vinculacion3.0",
			Hostname:        "190.122.229.53:3306",
			Database:        "VEDB",
			MaxIddleConn:    1,
			MaxOpenConn:     3,
			ConnMaxLifetime: time.Minute * 5,
		}
	/*} else {
		databaseConfig = DatabaseConfig{
			Username:        "root",
			Password:        "root",
			Hostname:        "localhost:3306",
			Database:        "biblio_digital",
			MaxIddleConn:    1,
			MaxOpenConn:     3,
			ConnMaxLifetime: time.Minute * 5,
		}
	}*/

	return databaseConfig
}

type table struct {
	Name  string
	Model interface{}
}

var (
	bibliotecaDigitalDB *gorm.DB
)

func init() {
	databaseConfig := getDataBaseSpecification()
	log.Println("init_db --> databaseConfig:", databaseConfig)

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&autoCommit=true", databaseConfig.Username, databaseConfig.Password, databaseConfig.Hostname, databaseConfig.Database)
	bibliotecaDigitalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("db_connection_error database:", databaseConfig.Database)
		panic(err)
	}
}

func BibliotecaDigitalDB() (*gorm.DB, error) {
	var err error

	db, err := bibliotecaDigitalDB.DB()
	stats := db.Stats()

	if stats.OpenConnections >= 10 {
		err = fmt.Errorf("number of connections exceeded: %v", stats.OpenConnections)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("error connect biblioteca_digital db", err)
	}

	return bibliotecaDigitalDB, err
}
