package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"russell/sql-test/config"
)

var dbConfString string

/**
 * @param dbType 设置的数据库类型
 */
func ConnDB(dbType string) (*gorm.DB) {

	confPath := "src/russell/sql-test/config/env.yaml"
	if len(os.Args) == 5 && os.Args[4] == "test" {
		confPath = "../config/env.yaml"
	}

	config.LoadConfiguration(confPath)

	var dbConf = config.GetConfiguration()
	var currDBConf DBConf

	SetDBConnectInfo(dbType, &currDBConf, *dbConf)

	// "user:password@(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConfString)
	if err != nil {
		log.Fatalln(err)
	}

	db.SingularTable(true)
	return db
}

/**
 * 根据dbType类型设置数据库链接配置
 */
func SetDBConnectInfo(dbType string, currDBConf *DBConf, dbConf config.Configuration) {
	switch dbType {
	case "core":
		SetCoreDBInfo(currDBConf, dbConf)
		break
	case "sku":
		SetSKUDBInfo(currDBConf, dbConf)
		break
	}

	dbConfString = currDBConf.User + ":" + currDBConf.Pass + "@(" + currDBConf.Host + ":" + currDBConf.Port + ")/" + currDBConf.Name + "?charset=utf8&parseTime=True&loc=Local"
}

/**
 * 设置为core数据库的连接配置
 */
func SetCoreDBInfo(currDBConf *DBConf, dbConf config.Configuration) {
	currDBConf.Host = dbConf.CORE_DB_HOST
	currDBConf.Port = dbConf.CORE_DB_PORT
	currDBConf.Name = dbConf.CORE_DB_NAME
	currDBConf.User = dbConf.CORE_DB_USER
	currDBConf.Pass = dbConf.CORE_DB_PASSWORD
}

/**
 * 设置为sku数据库的连接配置
 */
func SetSKUDBInfo(currDBConf *DBConf, dbConf config.Configuration) {
	currDBConf.Host = dbConf.SKU_DB_HOST
	currDBConf.Port = dbConf.SKU_DB_PORT
	currDBConf.Name = dbConf.SKU_DB_NAME
	currDBConf.User = dbConf.SKU_DB_USER
	currDBConf.Pass = dbConf.SKU_DB_PASSWORD
}
