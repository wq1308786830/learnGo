package biz_tests

import (
	"github.com/jinzhu/gorm"
	"log"
	"russell/sql-test/models"
)

var result *gorm.DB

func TestsRegister() {
	result = models.DeleteByOpenId("dsadas1dasfff4rfgdsddfg")
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
}
