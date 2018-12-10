package sql_test

import "russell/sql-test/biz-tests"

func StartTests() {
	openid := "123123213"
	biz_tests.Register(openid)
}
