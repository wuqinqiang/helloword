package db

import (
	_ "embed"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Provider interface {
	Init() error
	GetDb() (*gorm.DB, error)
}

func Init(provider Provider) error {
	err := provider.Init()
	if err != nil {
		return err
	}
	db, err = provider.GetDb()
	if err != nil {
		return err
	}
	return nil
}

func Get() *gorm.DB {
	if db == nil {
		panic("db must not nil,please call the sqlite first")
	}
	return db
}
