package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your text: ")
	base, _ := reader.ReadString('\n')
	toDouble := []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'z'}

	var valueMatch bool

	for _, value1 := range base {
		for _, value2 := range toDouble {
			if strings.ToLower(string(value1)) == string(value2) {
				valueMatch = true
			}
		}

		if valueMatch{
			fmt.Print(string(value1) + "o" + strings.ToLower(string(value1)))
		}else{
			fmt.Print(string(value1))
		}

		valueMatch = false
	}

}
