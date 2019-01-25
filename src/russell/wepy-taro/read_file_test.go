/*
官网教程
*/
package wepy_taro

import (
	"fmt"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"/Users/russell/Projects/saas-2c-weichatApp/src/page/line/newInfoForm1.wpy", "dlrow ,olleH"},
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
	path := "/Users/russell/Projects/saas-2c-weichatApp/src/page/line/newInfoForm1.wpy"
	contentByte := GetFileContent(path)
	temp := SplitWpyFile(contentByte)
	CreateFileByTemp(temp, path)
}
