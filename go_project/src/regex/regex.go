package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ccmouse@gmail.com@abc.com
email1 is abc@a.com
email2 is 123@a.com
emial3 is 456@a.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}
