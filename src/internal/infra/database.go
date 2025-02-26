package infra

import (
	"fmt"
	"github.com/ChristianSilvaDev/GoMail/src/internal/interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MySQLDatabase struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	TimeZone        string
	Entities        []*interfaces.Entity
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

var dbConnectionCache *gorm.DB = nil

func (mdb *MySQLDatabase) Init(entities ...interfaces.Entity) {
	conn, err := mdb.GetConnection()

	if err != nil {
		panic(err)
	}

	for _, entity := range entities {
		err = conn.AutoMigrate(entity)
		if err != nil {
			panic(err)
		}
	}
}

func (mdb *MySQLDatabase) GetConnection() (*gorm.DB, error) {
	if dbConnectionCache != nil {
		return dbConnectionCache, nil
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		mdb.User, mdb.Password, mdb.Host, mdb.Port, mdb.DBName, mdb.TimeZone,
	)

	fmt.Println(dsn)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := conn.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(mdb.MaxIdleConns)
	db.SetMaxOpenConns(mdb.MaxOpenConns)
	db.SetConnMaxLifetime(mdb.ConnMaxLifetime)

	dbConnectionCache = conn
	return conn, nil
}
