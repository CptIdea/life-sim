package generator

import "math/rand"

const symbols = " <>^v_|?#@:\\$1234567890+-/*%!`"

func GenerateCode(x, y int) string {
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
