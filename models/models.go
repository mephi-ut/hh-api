package models

//go:generate reform

import (
	cfg "github.com/mephi-ut/hh-api/config"
	db "github.com/mephi-ut/hh-api/db"
	"github.com/xaionaro/reform"
)

func init() {
	cfg.AddReloadHook(Reinit)
}

type modelI interface {
	SetDefaultDB(*reform.DB) error
	Table() reform.Table
}

func List() []modelI {
	return []modelI{
		&user{},
		&application{},
		&vacancy{},
	}
}

func Init() {
	db := db.Get()
	for _, model := range List() {
		model.Table().CreateTableIfNotExists(db)
		model.SetDefaultDB(db)
	}
}

func Reinit() {
	Init()
}
