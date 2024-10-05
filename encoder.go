package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"
)

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func numEncoder(num int) string {
	if num == 0 {
		return string(charset[0])
	}

	var encoded strings.Builder
	base := len(charset)

	for num > 0 {
		remainder := num % base
		encoded.WriteByte(charset[remainder])
		num = num / base
	}

	return reverse(encoded.String())
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}

func getUserInput(prompt string, reader *bufio.Reader) int {
	for {
		fmt.Println(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		return number
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	num := getUserInput("Enter a number to encode: ", reader)
	encoded := numEncoder(num)
	fmt.Println("base62 encoded: ", encoded)
}
