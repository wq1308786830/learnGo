package sql_test

func ApplicationStart() {

	StartTests()
	StartServer()
}


//
//func dbOp() {
//
//	if err != nil {
//		log.Fatalln("err open databases", err)
//	}
//
//	DB.AutoMigrate(&models.SassComponentRelation{})
//
//	var userRelation models.SassComponentRelation
//
//	DB.Where(&models.SassComponentRelation{UserId: "2"}).Find(&userRelation)
//
//	if userRelation.UserId == "" {
//		userRelation = models.SassComponentRelation{UserId: "2", WechatKhaosOpenid: "332323", WechatKhaosOpenidH5: "11111"}
//
//		fmt.Println(DB.NewRecord(userRelation))
//		fmt.Println(DB.Create(&userRelation))
//	}
//
//	DB.First(&userRelation, "wechat_khaos_openid = ?", "34")
//	fmt.Println("userID:", userRelation.UserId)
//}
//
//func dbDefer() {
//	var userRelation1 models.SassComponentRelation
//	DB.First(&userRelation1, "wechat_khaos_openid = ?", "dsadasdasfff4rfgdsddfg")
//	fmt.Println("userID1:", userRelation1.UserId)
//}

