package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {

	//	Membalikkan string
	var reverse string

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}

	//	Menambah tanda _
	var result string

	for i := 0; i < len(reverse); i++ {
		if i+1 < len(reverse) && reverse[i+1] == ' ' || reverse[i] == ' ' {
			result += string(reverse[i])
		} else {
			result += string(reverse[i]) + "_"
		}
	}

	return result[:len(result)-1] // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
	fmt.Println(ReverseString("I am a student"))
}
