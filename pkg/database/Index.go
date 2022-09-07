package database

import (
	"fmt"
	"strings"

	_ "gitee.com/travelliu/dm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"

	"xxqg-tk/pkg/config"
)

var DB *xorm.Engine

func init() {
	dbType := config.GetString("database.type")
	switch strings.ToLower(dbType) {
	case "sqlite":
		initSqlite()
	case "oracle":
		initOracle()
	case "dm":
		initDm()
	case "pq":
		fallthrough
	case "postgres":
		initPostgres()
	case "mysql":
		initMysql()
	default:
		logger.Fatal("暂未开通配置的数据库类型")
	}
	if err := DB.Ping(); err != nil {
		logger.Fatal("数据库连接失败! %v", err)
	}
	if config.GetString("database.prefix") != "" {
		DB.SetTableMapper(names.NewPrefixMapper(names.SnakeMapper{}, config.GetString("database.prefix")))
	}
	if config.GetBool("database.showSql") {
		DB.ShowSQL(true)
	}
}

func initSqlite() {
	var err error
	DB, err = xorm.NewEngine("sqlite3", config.GetString("database.address"))
	if err != nil {
		logger.Fatal("数据库创建失败! %v", err)
	}
}

func initOracle() {
	logger.Fatal("尚未支持")
}

func initDm() {
	var err error
	DB, err = xorm.NewEngine("dm", fmt.Sprintf("dm://%s:%s@%s:%d",
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.address"),
		config.GetInt("database.port"),
	))
	if err != nil {
		logger.Fatal("数据库创建失败! %v", err)
	}
}

func initPostgres() {
	var err error
	DB, err = xorm.NewEngine("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.GetString("database.address"),
		config.GetInt("database.port"),
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.dbName"),
		config.GetString("database.sslMode"),
	))
	if err != nil {
		logger.Fatal("数据库创建失败! %v", err)
	}
}

func initMysql() {
	var err error
	DB, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.address"),
		config.GetInt("database.port"),
		config.GetString("database.dbName"),
		config.GetString("database.charset"),
	))
	if err != nil {
		logger.Fatal("数据库创建失败! %v", err)
	}
}
