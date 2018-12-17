package biz_tests

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"net/url"
	"russell/sql-test/config"
	"russell/sql-test/models"
	"russell/sql-test/utils"
)

var result *gorm.DB
var env = config.GetConfiguration()

/**
用户注册模拟
 */
func Register(openid string) int {
	result = models.DeleteByOpenId(openid)
	if result.Error != nil {
		log.Println(result.Error)
		return utils.DBERRORCODE
	}

	return OpenLogin(openid)
}

/**
用户登录模拟
 */
func OpenLogin(openid string) int {

	urlStr := env.SERVICE_API_URL + env.KHAOS_VERSION + "/openlogin"
	param := url.Values{"openid": {openid}, "originType": {"miniprogram"}}

	resp := utils.HttpPost(urlStr, param)

	var data map[string]interface{}
	json.Unmarshal(resp, &data)

	var code int
	if data["code"] != nil {
		code = int(data["code"].(float64))
	} else {
		code = int(data["status"].(float64))
	}

	return code
}
