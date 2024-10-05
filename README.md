# BASE 62 ENCODER

### A README file TO HELP MY OLDER SELF UNDERSTAND THE TRASH CODE

### intro: 
Base 62 encoder is a endoding technique of encoding integer (numbers) into strings (A-Z) (a-z) (0-9)
```
26 uppercase
26 lowercase
10 numbers

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
```

### steps to encode a number

- divide the number by 62
- the remainder will give you corresponding Base62 character
- Repeat until the quotient becomes 0
- **Reverse** the result since the last remainder corresponds to the first character in Base62.

```
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

```

### steps to decode 

```
func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}
```

### getting user input
taking string output and int output

```
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
```

converting the string to ascii value
```
		number, err := strconv.Atoi(input)
```

### main function

```
func main() {
	reader := bufio.NewReader(os.Stdin)
	num := getUserInput("Enter a number to encode: ", reader)
	encoded := numEncoder(num)
	fmt.Println("base62 encoded: ", encoded)
}

```
