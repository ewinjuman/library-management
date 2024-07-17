package database

import (
	"fmt"
	Session "github.com/ewinjuman/go-lib/session"
	"library-management/CategoryService/app/domain/entities"
	"library-management/CategoryService/pkg/repository"
	"library-management/CategoryService/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var dbConnection *gorm.DB

func init() {
	err := mysqlOpen()
	if err != nil {
		panic(err.Error())
	}
	dbConnection.AutoMigrate(&entities.Category{})
}

// Mysql open connection
func mysqlOpen() error {
	//var err error
	//config := configs.Config.Database

	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.Open(mysqlConnURL))
	if err != nil {
		//fmt.Println("Failed to connect database. err: ", err.Error())
		return fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(5)

	dbConnection = db
	return nil
}

// MysqlConnection func for connection to Mysql database.
func MysqlConnection(session *Session.Session) (*gorm.DB, error) {
	if dbConnection == nil {
		if err := mysqlOpen(); err != nil {
			session.Error(err.Error())
			return dbConnection, repository.UndefinedErr
		}
	}
	sqlDB, err := dbConnection.DB()
	if err != nil {
		return nil, err
	}
	if errping := sqlDB.Ping(); errping != nil {
		errping = nil
		if errping = mysqlOpen(); errping != nil {
			session.Error(errping.Error())
			return dbConnection, repository.UndefinedErr
		}
	}
	logLevel := logger.Info
	//if !configs.Config.Database.LogMode {
	//	logLevel = logger.Silent
	//}
	newLogger := logger.New(
		session.Logger, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	//dbConnection.Logger.LogMode(logger.Silent)
	dbConnection.Logger = newLogger
	return dbConnection, nil
}
