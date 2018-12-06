package main

import (
	"fmt"
	"log"
	"russell/sql-test"
	"russell/sql-test/models"
)

func main() {
	db, err := sql_test.ConnDB("core")
	if err != nil {
		log.Fatalln("err open databases", err)
	}

	db.AutoMigrate(&models.SassComponentRelation{})

	var userRelation models.SassComponentRelation

	db.Where(&models.SassComponentRelation{UserId: "3"}).Find(&userRelation)

	if userRelation.UserId == "" {
		userRelation = models.SassComponentRelation{UserId: "2", WechatKhaosOpenid: "332323", WechatKhaosOpenidH5: "11111"}

		fmt.Println(db.NewRecord(userRelation))
		fmt.Println(db.Create(&userRelation))
	}

	db.First(&userRelation, "wechat_khaos_openid = ?", "332323")
	log.Fatalln("userID:", userRelation.UserId)
}
