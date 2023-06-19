package main

import (
	"fmt"
	"os"
	"strings"
)

func ConvertStr(str, banner string) string {

	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	dataSplit := strings.Split(string(data), "\n")
	strTab := strings.Split(str, "\r\n")
	var newStr string
	for j, args := range strTab {
		if args != "" {
			var result []string
			for _, char := range args {
				for i := 1; i <= 8; i++ {
					result = append(result, dataSplit[((char-32)*9)+rune(i)])
				}
			}
			var tab [8][]string
			for i, val := range result {
				tab[i%8] = append(tab[i%8], val)
			}
			for _, ligne := range tab {
				for _, part := range ligne {
					newStr += part
				}
				newStr += "\n"

			}
		} else if j != len(str)-1 && args != "\r" {
			newStr += "\n"
		}
	}
	return newStr
}


func isValid(s string) bool {
	str :=[]rune(s)
	for _, ch := range str {
		if ch < ' ' || ch != 10 || ch != '\r' || ch > '~' {
			return false
		}
	}
	return true
}
func isPrintable(s string) bool{
	str:=[]rune(s)
	for _,ch := range str {
		if ch>=32 && ch <=125{
			return true
		}
		
	}
	return false
}
