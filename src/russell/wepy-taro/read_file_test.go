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
	var dst []byte
	templateStr := []byte(`$1|$2|$3|$4|$5|$6`)
	xmlContent := GetFileContent("/Users/russell/Desktop/wpysrc/page/line/newInfoForm1.xml")
	//fmt.Println(string(xmlContent))
	regPropsReact := regexp.MustCompile(`(?ms)(<\w+)(?U: className=".*")(?U: wx:if={ ?| wx:elif={ ?| wx:else)([ !&|.\w]+)(?U: ?})(.*)`) // 匹配所有组件传值属性
	match := regPropsReact.FindAllSubmatchIndex(xmlContent, -1)
	for _, v := range match {
		fmt.Printf("%q\n", regPropsReact.Expand(dst, templateStr, xmlContent, v))
	}
}
