package generator

import "math/rand"

const symbols = " <>^v_|?#@:\\$1234567890+-/*%!`"

//const start = " $9!05%%! \n2%>>^?1/>5\n^!^%^>5821\n5>  /@/9#^\n!1^8@^?19\\\n/ 45>6\\?>2\n$^* */++68\n?2 $*20_6^\n268%%#1$/<\n- \\9!!<v6>"
const start = ""

func GenerateCode(x, y int) string {
	if start != "" {
		return start
	}
	ans := ""
	for i := 0; i < y; i++ {
		ans += generateLine(x) + "\n"
	}
	return ans
}

func generateLine(len int) string {
	ans := ""
	for i := 0; i < len; i++ {
		ans += generateChar()
	}
	return ans
}

func generateChar() string {
	return string(symbols[rand.Intn(len(symbols)-1)])
}
