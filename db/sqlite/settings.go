package sqlite

import (
	"context"
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"

	"github.com/pkg/errors"
	"github.com/wuqinqiang/helloword/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	//go:embed helloword.sql
	execQql string
)

type Settings struct {
	//the db saved path
	path string
	// db fileName
	dbFileName string
	// sql for sqlite
	execSql      string
	maxLifetime  int
	maxIdleConns int
	maxOpenConns int
	// logger
	logger logger.Interface
}

func New(path string) *Settings {
	settings := DefaultSettings
	if path == "" {
		path = "~/.bridge"
	}
	settings.path = path
	return settings
}

func (settings *Settings) Init() error {
	return settings.init()
}

func (settings *Settings) init() error {
	// if default path have '~'
	actualPath, err := homedir.Expand(settings.path)
	if err != nil {
		return err
	}
	settings.path = actualPath
	_, err = os.Stat(filepath.Join(settings.path, settings.dbFileName))
	noexist := os.IsNotExist(err)
	if err != nil && !noexist {
		return err
	}
	if noexist {
		err = os.MkdirAll(settings.path, 0755) //nolint: gosec
		if err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}

func (settings *Settings) GetDb() (*gorm.DB, error) {
	gormDb, err := gorm.Open(sqlite.Open(filepath.Join(settings.path, settings.dbFileName+"?cache=shared")), &gorm.Config{
		Logger: settings.logger,
	})
	if err != nil {
		panic(err)
	}
	sqlDb, err := gormDb.DB()
	if err != nil {
		return nil, err
	}

	if settings.execSql != "" {
		_, err = sqlDb.ExecContext(context.Background(), settings.execSql)
		if err != nil {
			//hard code for duplicate column
			if !strings.Contains(err.Error(), "duplicate column name") {
				return nil, errors.Wrap(err, "ExecContext")
			}
			logging.Warnf("[ExecContext] err:%v", err)
		}
	}
	sqlDb.SetMaxOpenConns(settings.maxOpenConns)
	sqlDb.SetMaxIdleConns(settings.maxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(settings.maxLifetime) * time.Second)
	return gormDb, nil
}
