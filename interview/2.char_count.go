package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	Left = iota
	Top
	Right
	Bottom
)

func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func reverString(s string) (string, bool) {
	str := []byte(s)
	//str := s
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
}

func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}

func move(cmd string, x0 int, y0 int, z0 int) (x, y, z int) {
	x, y, z = x0, y0, z0
	repeat := 0
	repeatCmd := ""
	for _, s := range cmd {
		switch {
		case unicode.IsNumber(s):
			repeat = repeat*10 + (int(s) - '0')
		case s == ')':
			for i := 0; i < repeat; i++ {
				x, y, z = move(repeatCmd, x, y, z)
			}
			repeat = 0
			repeatCmd = ""
		case repeat > 0 && s != '(' && s != ')':
			repeatCmd = repeatCmd + string(s)
		case s == 'L':
			z = (z + 1) % 4
		case s == 'R':
			z = (z - 1 + 4) % 4
		case s == 'F':
			switch {
			case z == Left || z == Right:
				x = x - z + 1
			case z == Top || z == Bottom:
				y = y - z + 2
			}
		case s == 'B':
			switch {
			case z == Left || z == Right:
				x = x + z - 1
			case z == Top || z == Bottom:
				y = y + z - 2
			}
		}
	}
	return
}

func match() {
	buf := "abc azc a7c aac 888 a9c  tac"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

// find    \d*\([LFRB]*\)
// than open it

func match2() string {
	buf := "R2(LF)2(RF)B"
	fmt.Println(buf)
	//解析正则表达式，如果成功返回解释器
	//reg1 := regexp.MustCompile(`\d*\(\w*\)`)
	reg1 := regexp.MustCompile(`\d*\([LFRB]*\)`)
	//reg1 := regexp.MustCompile(`\(.*\)`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return ""
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	//fmt.Println("result1 = ", result1)
	for _, x := range result1 {
		rep := open(x[0])
		buf = strings.Replace(buf, x[0], rep, 1)
	}
	return buf
}
func open(x string) string {
	result := regexp.MustCompile(`^\d*`)
	fact := result.FindString(x)
	result2 := regexp.MustCompile(`\([LFRB]*\)`)
	s := result2.FindString(x)
	s = strings.TrimLeft(s, "(")
	s = strings.TrimRight(s, ")")
	f, _ := strconv.Atoi(fact)
	result3 := strings.Repeat(s, f)
	return result3
}

func main() {

	s := "ulookthirt   y"
	s2 := "fuck"
	s2 = strings.Repeat(s2, 10)
	fmt.Println(isUniqueString(s))
	fmt.Println(s2)
	fmt.Println(isUniqueString(s2))

	fmt.Println(reverString(s))
	fmt.Println(replaceBlank(s))

	println(move("R2(LF)", 0, 0, Top))
	fmt.Println(match2())

}
