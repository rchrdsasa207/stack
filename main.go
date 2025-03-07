package main

import (
	"fmt"
	"strconv"
	"strings"
)

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func rpn(expr []string) int {
	var stack []int
	for _, v := range expr {
		switch v {
		case "+":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case "-":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
		case "/":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
		case "*":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		default:
			stack = append(stack, atoi(v))
		}
	}
	return stack[0]
}

func simplify(dir string) string {
	a := strings.Split(dir, "/")
	var path []string
	for _, v := range a {
		switch v {
		case ".":
			continue
		case "":
			continue
		case "..":
			if len(path) != 0 {
				path = path[:len(path)-1]
			}
		default:
			path = append(path, v)
		}
	}
	return "/" + strings.Join(path, "/")
}

/*
	func npn(in string) (success bool, num int, operation string) {
		if len(in) == 0 {
			return false, 0, ""
		}
		suc, num, opoperation := npn(in[1:])
		switch v {
		case "+":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case "-":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
		case "/":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
		case "*":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		default:
			stack = append(stack, atoi(v))
		}
	}
*/
func main() {
	fmt.Println(simplify("/foo/./bar/../../baz/"))
}
