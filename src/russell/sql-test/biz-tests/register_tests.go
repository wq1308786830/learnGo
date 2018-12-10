package biz_tests

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"russell/sql-test/config"
	"russell/sql-test/models"
	"russell/sql-test/utils"
)

var result *gorm.DB

func Register(openid string) int {
	result = models.DeleteByOpenId(openid)
	if result.Error != nil {
		log.Println(result.Error)
		return utils.DBERRORCODE
	}
	return OpenLogin(openid)
}

func OpenLogin(openid string) int {
	env := config.GetConfiguration()
	url := env.SERVICE_API_URL + env.KHAOS_VERSION + "/openlogin"
	param := make(map[string]interface{})

	param["openid"] = openid
	param["originType"] = "miniprogram"
	params, err := json.Marshal(param)

	if err != nil {
		log.Println(err)
	}

	resp, err := utils.HttpPost(url, params)

	var data map[string]interface{}
	json.Unmarshal(resp, &data)

	var code int
	if data["code"] != nil {
		code = data["code"].(int)
	} else {
		code = int(data["status"].(float64))
	}

	return code
}
