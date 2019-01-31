package wepy_taro

import (
	"bytes"
	"fmt"
	"regexp"
)

func ReplaceBaseXml(xmlContent []byte) []byte { // 匹配所有tag的class属性
	var regViewS = regexp.MustCompile(`<(\w)|</(\w)`)                        // 匹配所有tag前两位，分组匹配tag第一个字母
	var regComment = regexp.MustCompile(` class="|<!--|-->|"{{|}}"|{{|}}`)   // 匹配所有xml注释以及 class="
	var regProps = regexp.MustCompile(`(?U)([ :]?[.\w-_@]+[.sync]?)=(".+")`) // 匹配所有组件传值属性

	// 替换匹配所有tag前两位，分组匹配tag第一个字母
	xmlContent = regViewS.ReplaceAllFunc(xmlContent, func(s []byte) []byte {
		return bytes.ToUpper(s)
	})

	// 替换匹配所有xml注释以及 class="
	xmlContent = regComment.ReplaceAllFunc(xmlContent, func(i []byte) []byte {
		var result []byte
		switch string(i) {
		case ` class="`:
			result = []byte(" className=\"")
		case `<!--`:
			result = []byte("{/* <!--")
		case `-->`:
			result = []byte("--> */}")
		case `"{{`, `{{`:
			result = []byte("{")
		case `}}"`, `}}`:
			result = []byte("}")
		default:
			result = i
		}
		return result
	})

	//替换匹配所有组件传值属性
	xmlContent = regProps.ReplaceAllFunc(xmlContent, func(s []byte) []byte {
		// 同Taro属性以及值不用转换
		regReactProps := regexp.MustCompile(`className="|type="|placeholder="|range="|value="|formType="|mode="|src="`)
		classNameLen := regReactProps.Find(s)
		if len(classNameLen) > 0 {
			//fmt.Println("if", string(s))
			return s
		}

		regWpyProps := regexp.MustCompile(`bindsubmit="|report-submit="|@blur="|@tap="|bindtap="|bindfocus="|bindchange="|bindinput="|@input="|wx:if="|placeholder-class="`)
		wpyProps := regWpyProps.FindAll(s, -1)
		if len(wpyProps) > 0 {
			if wpyProps == nil {
				return s
			}

			var result []byte
			var isEvent = true
			for _, value := range wpyProps {
				wpyPropsStr := string(value)
				switch wpyPropsStr {
				case `bindsubmit="`:
					result = []byte(`onSubmit="`)
				case `report-submit="`:
					isEvent = false
					result = []byte(`reportSubmit="`)
				case `@blur="`, `bindblur="`:
					result = []byte(`onBlur="`)
				case `@tap="`, `bindtap="`:
					result = []byte(`onClick="`)
				case `@focus="`, `bindfocus="`:
					result = []byte(`onFocus="`)
				case `@change="`, `bindchange="`, `@input="`, `bindinput="`:
					result = []byte(`onChange="`)
				case `placeholder-class="`:
					isEvent = false
					result = []byte(`placeholderClass="`)
				default:
					isEvent = false
					result = value
				}

				result = regWpyProps.ReplaceAll(s, result)

				if isEvent {
					fmt.Println(string(result))
					regEventFuncWithBrackets := regexp.MustCompile(`(=")([\w]+)(\()([0-9a-zA-Z-_"', ]+)(\)")`) // 匹配带参事件使用的方法
					result = regEventFuncWithBrackets.ReplaceAll(result, []byte("={this.$2.bind(this, $4)}"))

					regEventFunc := regexp.MustCompile(`(=")([^(][\w]+)(")`) // 匹配事件使用的方法
					result = regEventFunc.ReplaceAll(result, []byte("={this.$2}"))
					fmt.Println(string(result))
				} else {
					fmt.Println(string(result))
					regPropsVal := regexp.MustCompile(`=("{{)(.*)(}}")`) // 匹配属性值
					result = regPropsVal.ReplaceAll(result, []byte("={$2}"))
					fmt.Println(string(result))
				}
				//fmt.Println("isEvent", string(result))
				return result
			}
			return regWpyProps.ReplaceAll(s, result)
		} else {
			fmt.Println(string(s))
			regPropsReact := regexp.MustCompile(`(?::?)([\w_]+)(?:\.sync="|=")([\w_]+)"`) // 匹配所有组件传值属性
			s = regPropsReact.ReplaceAll(s, []byte(`$1={$2}`))
			fmt.Println(string(s))
			return s
		}
	})

	// 替换wx:if wx:else wx:elif
	//regIfElse := regexp.MustCompile(`(<\w+)(?U: wx:if={ ?| wx:elif={ ?| wx:else)([ !&|.\w]+)(?U: ?})(.*)`) // 匹配所有条件渲染
	//xmlContent = regIfElse.ReplaceAll(xmlContent, []byte(`{$2 && $1$3}`))

	return xmlContent
}

func ReplaceStatement(xmlContent []byte) []byte {

	return nil
}
