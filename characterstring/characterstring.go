package characterstring

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"regexp"
	"strings"
)

//反转字符串
func ReverseString(s []byte) {
	length := len(s)
	var aa byte
	for i := 0; i < length/2; i++ {
		aa = s[i]
		s[i] = s[length-i-1]
		s[length-i-1] = aa
	}
	fmt.Println(string(s))
}

//整数反转
func Reverse(x int) int {
	remainder := 0
	newValue := 0
	for x != 0 {
		remainder = x % 10
		x = x / 10
		newValue = newValue*10 + remainder
	}
	if newValue > math.MaxInt32 || newValue < math.MinInt32 {
		return 0
	}
	return newValue
}

//字符串中的第一个唯一字符
func FirstUniqChar(s string) int {
	for i := 0; i < len(s); i++ {
		flag := true
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}
			if s[i] == s[j] {
				flag = false
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

//有效的字母异位词
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_array := make(map[byte]int, 26)
	t_array := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		s_array[s[i]] += 1
		t_array[s[i]] += 1
	}
	return reflect.DeepEqual(s_array, t_array)
}

//验证回文串(正反读都一样)
func IsPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, "")
	length := len(s)

	for i := 0; i < length/2; i++ {
		a := string(s[i])
		b := string(s[length-1-i])
		if !strings.EqualFold(a, b) {
			return false
		}
	}
	return true
}

//字符串转换整数(atoi)
func MyAtoi(s string) int {
	index := 0
	length := len(s)
	for _, v := range s {
		if v == 32 {
			index++
		} else {
			break
		}
	}
	if index >= length {
		return 0
	}
	sign := 1
	if string(s[index]) == "+" || string(s[index]) == "-" {
		if string(s[index]) == "-" {
			sign = -1
		}
		index++
	}
	result := 0
	temp := 0
	for index < length {
		num := s[index] - 48
		if s[index] < 48 || s[index] > 57 {
			break
		}
		temp = result
		result = result*10 + int(num)
		if result/10 != temp || result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		index++
	}
	return result * sign
}

//实现strStr()
func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 外观数列

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	result := ""
	index := 0
	Minlength := len(strs[0])
	for _, s := range strs {
		if Minlength >= len(s) {
			Minlength = len(s)
		}
	}
	for index < Minlength {
		aa := strs[0][index]
		for _, s := range strs {
			if aa != s[index] {
				return result
			}
		}
		result += string(aa)
		index++
	}
	return result
}
