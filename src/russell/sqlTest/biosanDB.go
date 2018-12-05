package sqlTest

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"russell/sqlTest/models"
)

func ConnDB(dbType string) *gorm.DB {

	var dbConf Configuration
	var currDBConf models.DBConf

	switch dbType {
	case "core":
		currDBConf.Host = dbConf.CORE_DB_HOST
		currDBConf.Port = dbConf.CORE_DB_PORT
		currDBConf.Name = dbConf.CORE_DB_NAME
		currDBConf.User = dbConf.CORE_DB_USER
		currDBConf.Pass = dbConf.CORE_DB_PASSWORD
		break
	case "sku":
		currDBConf.Host = dbConf.SKU_DB_HOST
		currDBConf.Port = dbConf.SKU_DB_PORT
		currDBConf.Name = dbConf.SKU_DB_NAME
		currDBConf.User = dbConf.SKU_DB_USER
		currDBConf.Pass = dbConf.SKU_DB_PASSWORD
		break
	}

	db, err := gorm.Open("mysql", currDBConf.User+":"+currDBConf.Pass+"@"+currDBConf.Host+":"+currDBConf.Port+"/"+currDBConf.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	return db
}
