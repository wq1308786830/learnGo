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
		{"23324235555555", utils.SUCCESSCODE},
	}

	for _, c := range cases {
		got := biz_tests.Register(c.in)
		if got != c.want {
			t.Errorf("openid(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
