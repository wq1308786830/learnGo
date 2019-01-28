package wepy_taro

import (
	"bytes"
	"fmt"
	"regexp"
)

func ReplaceBase(xmlContent []byte) []byte {
	var regViewS = regexp.MustCompile(`<(\w)|</(\w)`)          // 匹配所有tag前两位，分组匹配tag第一个字母
	var regClass = regexp.MustCompile(` class="`)              // 匹配所有tag的class属性
	var regProps = regexp.MustCompile(`(:?.+[.sync]?)="(.+)"`) // 匹配所有组件传值属性

	//regViewS.ReplaceAllString(xmlContent, "${1}")

	xmlContent = regViewS.ReplaceAllFunc(xmlContent, func(s []byte) []byte {
		return bytes.ToUpper(s)
	})

	templateClass := []byte(" className=\"")

	xmlContent = regClass.ReplaceAll(xmlContent, templateClass)

	xmlContent = regProps.ReplaceAllFunc(xmlContent, func(s []byte) []byte {
		fmt.Println(string(s))
		return []byte("a={a}")
	})

	return xmlContent
}
