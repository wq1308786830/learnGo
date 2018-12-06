package sql_test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"russell/sql-test/models"
)

var DB *gorm.DB
var dbConfString string

func ConnDB(dbType string) (*gorm.DB, error) {

	LoadConfiguration("src/russell/sql-test/env.yaml")

	var dbConf = GetConfiguration()
	var currDBConf models.DBConf

	SetDBConnectInfo(dbType, &currDBConf, *dbConf)

	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConfString)
	if err != nil {
		log.Fatalln(err)
	}

	db.SingularTable(true)
	return db, err
}

func SetDBConnectInfo(dbType string, currDBConf *models.DBConf, dbConf Configuration) {
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

func SetCoreDBInfo(currDBConf *models.DBConf, dbConf Configuration) {
	currDBConf.Host = dbConf.CORE_DB_HOST
	currDBConf.Port = dbConf.CORE_DB_PORT
	currDBConf.Name = dbConf.CORE_DB_NAME
	currDBConf.User = dbConf.CORE_DB_USER
	currDBConf.Pass = dbConf.CORE_DB_PASSWORD
}

func SetSKUDBInfo(currDBConf *models.DBConf, dbConf Configuration) {
	currDBConf.Host = dbConf.SKU_DB_HOST
	currDBConf.Port = dbConf.SKU_DB_PORT
	currDBConf.Name = dbConf.SKU_DB_NAME
	currDBConf.User = dbConf.SKU_DB_USER
	currDBConf.Pass = dbConf.SKU_DB_PASSWORD
}
