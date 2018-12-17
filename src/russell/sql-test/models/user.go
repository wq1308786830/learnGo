package models

import "github.com/jinzhu/gorm"

var db = ConnDB("core")

type SassComponentRelation struct {
	UserId              string
	WechatKhaosOpenid   string
	WechatKhaosOpenidH5 string
}

func DeleteByOpenId(openid string) *gorm.DB {
	return db.Delete(SassComponentRelation{}, "wechat_khaos_openid = ?", openid)
}
