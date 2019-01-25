package wepy_taro

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
)

func GetFileContent(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	contentByte, readErr := ioutil.ReadAll(f)
	if readErr != nil {
		fmt.Println(readErr)
		return nil
	}

	return contentByte
}

/**
将wpy文件分解为less，xml，script内容
*/
func SplitWpyFile(content []byte) ContentTemp {
	styleReg := regexp.MustCompile(`(?ms)^(?-ms:<style.*>)(.*)(?-m:</style>)$`)
	xmlReg := regexp.MustCompile(`(?ms)^(?-ms:<template.*>)(.*)(?-m:</template>)$`)
	scriptReg := regexp.MustCompile(`(?ms)^(?-ms:<script.*>)(.*)(?-m:</script>)$`)
	contentStr := string(content)

	//fmt.Println(contentStr)
	lessStr := styleReg.FindStringSubmatch(contentStr)[1]
	xmlStr := xmlReg.FindStringSubmatch(contentStr)[1]
	scriptStr := scriptReg.FindStringSubmatch(contentStr)[1]

	contentTemp := ContentTemp{
		LessStr:   []byte(lessStr),
		XmlStr:    []byte(xmlStr),
		ScriptStr: []byte(scriptStr),
	}

	return contentTemp
}

/**
通过SplitWpyFile返回的content生成文件
*/
func CreateFileByTemp(content ContentTemp, path string) {

	var f *os.File
	var err error
	var realPath string // 实际文件路径
	var bytes []byte    // 实际写入对应文件的内容

	t := reflect.TypeOf(content)
	reg := regexp.MustCompile(`(\.wpy$)`) //匹配wpy文件正则

	for k := 0; k < t.NumField(); k++ {
		switch t.Field(k).Name {
		case "LessStr":
			bytes = content.LessStr
			realPath = reg.ReplaceAllString(path, ".less")
		case "XmlStr":
			bytes = content.XmlStr
			realPath = reg.ReplaceAllString(path, ".xml")
		case "ScriptStr":
			bytes = content.ScriptStr
			realPath = reg.ReplaceAllString(path, ".js")
		}

		f, err = os.Create(realPath)
		if err != nil {
			panic(err)
		}
		f.Write(bytes)

		fmt.Printf("%s \n", t.Field(k).Name)
	}
	defer f.Close()
}
