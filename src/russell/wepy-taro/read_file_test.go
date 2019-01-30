/*
官网教程
*/
package wepy_taro

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"/Users/russell/Desktop/wpysrc/page/line/newInfoForm1.wpy", "dlrow ,olleH"},
		{"/Users/russell/GitHub/GoLang/learn/src/russell/stringutil/reverse.go", "界世 ,olleH"},
	}

	for _, c := range cases {
		got := GetFileContent(c.in)
		fmt.Println(string(got))
		//if got != c.want {
		//	t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		//}
	}
}

func TestSplitWpyFile(t *testing.T) {
	path := "/Users/russell/Desktop/wpysrc/page/line/newInfoForm1.wpy"
	contentByte := GetFileContent(path)
	temp := SplitWpyFile(contentByte)
	CreateFileByTemp(temp, path)
}

func TestReg(t *testing.T) {
	s := []byte(`:imgData.sync="consentData"`)
	var dst []byte
	fmt.Println(string(s))
	regPropsReact := regexp.MustCompile(`(:?)([\w_]+)([.sync]?)(=")([\w_]+)(")`) // 匹配所有组件传值属性
	template := []byte("|$1|$2|$3|$4|$5|$6|")
	match := regPropsReact.FindSubmatchIndex(s)
	fmt.Printf("%q", regPropsReact.Expand(dst, template, s, match))
	fmt.Println(string(dst))
}
