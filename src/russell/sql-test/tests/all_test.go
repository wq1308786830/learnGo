package tests_test

import (
	"russell/sql-test/biz-tests"
	"russell/sql-test/utils"
	"testing"
)

func TestRegister(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"23123123213", utils.SUCCESSCODE},
		{"34", utils.SUCCESSCODE},
	}

	for _, c := range cases {
		got := biz_tests.Register(c.in)
		if got != c.want {
			t.Errorf("openid(%q) 结果 %d, 预期 %d", c.in, got, c.want)
		}
	}
}

func TestOpenLogin(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"osi8b0fM1Vx_6QJHBHQ8NIGubl78", utils.SUCCESSCODE},
		{"osi8b0faR-WY_WUuL7U2cQS2laEw", utils.SUCCESSCODE},
	}

	for _, c := range cases {
		got := biz_tests.OpenLogin(c.in)
		if got != c.want {
			t.Errorf("openid(%q) 结果 %d, 预期 %d", c.in, got, c.want)
		}
	}
}

func TestGetProjects(t *testing.T) {

}

