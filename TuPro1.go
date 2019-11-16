package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type ar [99] int

func state_pqrs() int {return 1}
func state_bukakurung() int {return 9}
func state_tutupkurung() int {return 10}
func state_error() int {return -1}
func state_n(c string) int {
	if c[1] == 'o' {
		return state_no(c)
	} else {
		return state_error()
	}
}
func state_no(c string) int {
	if c[2] == 't' {
		return state_not(c)
	} else {
		return state_error()
	}
}
func state_not(c string) int {
	if len(c) > 3 {
		return state_error()
	} else {
		return 2
	}
}
func state_a(c string) int {
	if c[1] == 'n' {
		return state_an(c)
	} else {
		return state_error()
	}
}
func state_an(c string) int {
	if c[2] == 'd' {
		return state_and(c)
	} else {
		return state_error()
	}
}
func state_and(c string) int {
	if len(c) > 3 {
		return state_error()
	} else {
		return 3
	}
}
func state_o(c string) int {
	if c[1] == 'r' {
		return state_or(c)
	} else {
		return state_error()
	}
}
func state_or(c string) int {
	if len(c) > 2 {
		return state_error()
	} else {
		return 4
	}
}
func state_x(c string) int {
	if c[1] == 'o' {
		return state_xo(c)
	} else {
		return state_error()
	}
}
func state_xo(c string) int {
	if c[2] == 'r' {
		return state_xor(c)
	} else {
		return state_error()
	}
}
func state_xor(c string) int {
	if len(c) > 3 {
		return state_error()
	} else {
		return 5
	}
}
func state_i(c string) int {
	if c[1] == 'f' {
		return state_if(c)
	} else {
		return state_error()
	}
}
func state_if(c string) int {
	if len(c) > 2 {
		if c[2] == 'f' {
			return state_iff(c)
		} else {
			return state_error()
		}
	} else {
		return 6
	}
}
func state_iff(c string) int {
	if len(c) > 3 {
		return state_error()
	} else {
		return 8
	}
}
func state_t(c string) int {
	if c[1] == 'h' {
		return state_th(c)
	} else {
		return state_error()
	}
}
func state_th(c string) int {
	if c[2] == 'e' {
		return state_the(c)
	} else {
		return state_error()
	}
}
func state_the(c string) int {
	if c[2] == 'e' {
		return state_then(c)
	} else {
		return state_error()
	}
}
func state_then(c string) int {
	if len(c) > 4 {
		return state_error()
	} else {
		return 7
	}
}

func getToken(c string) int {
	t := -1
	x := rune(c[0])
	if isProps(x) {
		t = state_pqrs()
	} else if x == '(' {
		t = state_bukakurung()
	} else if x == ')' {
		t = state_tutupkurung()
	} else if x == 'n' {
		t = state_n(c)
	} else if x == 'a' {
		t = state_a(c)
	} else if x == 'o' {
		t = state_o(c)
	} else if x == 'x' {
		t = state_x(c)
	} else if x == 'i' {
		t = state_i(c)
	} else if x == 't' {
		t = state_t(c)
	}
	return t
}

func isProps(c rune) bool {
	props := []rune{112, 113, 114, 115}
	for i := 0; i < len(props); i++ {
		if props[i] == c {
			return true
		}
	}
	return false
}

func lexer(T *ar, N *int, str string) {
	var token int
	word := ""
	i := 0
	for j, char := range str {
		_ = j
		if char != ' ' {
			if char == '(' || char == ')' {
				token = getToken(string(char))
			} else {
				word += string(char)
			}
			if j + 1 == len(str) && word != "" {
				token = getToken(word)
			}
		} else {
			if word != "" {
				token = getToken(word)
			} else {
				continue
			}
		}
		if token != 0 {
			T[i] = token
			i++
			*N = i
			if token == -1 {
				break
			}
			token = 0
			word = ""
		}
	}
}

func main() {

	var (
		token ar
		N int
	)

	fmt.Print("Input : ")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    str := scanner.Text()

	lexer(&token, &N, strings.ToLower(str))

	fmt.Print("Output: ")

	for i := 0; i < N; i++ {
		if token[i] == -1 {
			fmt.Print("error ")
		} else {
			fmt.Print(token[i], " ")
		}
	}
	fmt.Println()

}