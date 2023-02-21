package sqlite

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

var DefaultSettings = &Settings{
	maxLifetime:  7200,
	maxIdleConns: 5,
	maxOpenConns: 15,
	dbFileName:   "helloword.db",
	execSql:      execQql,
	logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	),
}
