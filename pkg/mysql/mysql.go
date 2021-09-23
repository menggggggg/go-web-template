package mysql

import (
	"database/sql"
	"fmt"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/menggggggg/go-web-template/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// New Create gorm.DB instance
func New(c *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	cfg, err := mysqlDriver.ParseDSN(c.DSN)
	if err != nil {
		return nil, err
	}

	err = createDatabaseWithMySQL(cfg)
	if err != nil {
		return nil, err
	}

	dialector = mysql.Open(c.DSN)

	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	logger.Infof("connect mysql success")
	return db, nil
}

func createDatabaseWithMySQL(cfg *mysqlDriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}
