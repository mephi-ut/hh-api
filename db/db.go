package db

import (
	"database/sql"
	"fmt"
	cfg "github.com/mephi-ut/hh-api/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xaionaro/reform"
	"github.com/xaionaro/reform/dialects/mysql"
	"github.com/xaionaro/reform/dialects/sqlite3"
	"time"
)

var db *reform.DB

var connectHooks []func()

func init() {
	cfg.AddReloadHook(Reinit)
}

type InitDBParams struct {
	cfg.DbCfg
}

func initDB(params InitDBParams) (db *reform.DB) {
	var connectionString string

	switch params.Driver {
	case "mysql":
		connectionString = getMysqlConnectionString(params)
	default:
		panic(fmt.Errorf(`Unknown DB driver: %v`, params.Driver))
		return nil
	}

	db = initDbByConnectionString(params, connectionString)

	return
}

func Init() {
	db = initDB(InitDBParams{DbCfg: cfg.Get().Db})
	if db == nil {
		panic(fmt.Errorf("Cannot connect to DB"))
	}
	for _, connectHook := range connectHooks {
		connectHook()
	}
}

func Reinit() {
	Init()
}

func Get() *reform.DB {
	return db
}

func printf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func initDbByConnectionString(params InitDBParams, connectionString string) *reform.DB {
	db, err := sql.Open(params.Driver, connectionString)
	if err != nil {
		panic(err)
	}

	setupDb(db, params.Driver)

	logger := smartLogger{dbName: params.Db, traceLogger: reform.NewPrintfLogger(printf), errorLogger: reform.NewPrintfLogger(printf)}
	logger.SetTraceEnable(true)
	logger.SetErrorEnable(true)

	switch params.Driver {
	case "mysql":
		return reform.NewDB(db, mysql.Dialect, logger)
	case "sqlite3":
		return reform.NewDB(db, sqlite3.Dialect, logger)
	}
	panic(fmt.Errorf("Unknown driver: ", params.Driver))
	return nil
}

func setupDb(db *sql.DB, driver string) {
	switch driver {
	case "sqlite3":
		db.SetMaxIdleConns(1)
		db.SetMaxOpenConns(1)
	case "mysql":
		db.Exec("SET wait_timeout=15")
		db.Exec("SET interactive_timeout=15")
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(1 * time.Minute)
		break
	default:
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		break
	}
}

func AddConnectHook(hook func()) {
	connectHooks = append(connectHooks, hook)
}
