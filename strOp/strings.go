package strop

import (
	"fmt"
	"strings"
)

// HasPreAndSuff - 判断是否以 ** 前缀 or 后缀 结尾
func HasPreAndSuff() {
	str := "This is an example of string"

	// 判断前缀
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")
	fmt.Print(" ")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	// 判断后缀
	fmt.Printf("T/F? Does the string \"%s\" have Suffix %s?", str, "ing")
	fmt.Print(" ")
	fmt.Printf("%t\n", strings.HasSuffix(str, "asd"))
}

// HasContain - 是否包含
func HasContain() {
	str := "This is a string"
	fmt.Printf("T/F? Does the string \"%s\" contain %s?", str, "is")
	fmt.Print(" ")
	fmt.Printf("%t\n", strings.Contains(str, "is"))
}

// FindIndexOfCharacter - 查找 子字符 or 子字符串是否在 当前字符串中
func FindIndexOfCharacter() {
	str := "This is a string"
	s := "i"
	fmt.Printf("The index of %s is %d", s, strings.Index(s, str))
}
