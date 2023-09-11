package main

import (
	"fibonacci/handler"
	"strconv"
	"strings"
)

func init() {
	a := HostImpl{}
	handler.SetHandler(a)
}

type HostImpl struct {
}

func fibonacciRange(m, n int) []int {
	if n < m || m < 0 || n < 0 {
		return nil
	}

	fib := make([]int, n-m+1)
	a, b := 0, 1

	for i := 0; i <= n; i++ {
		if i >= m {
			fib[i-m] = a
		}
		a, b = b, a+b
	}

	return fib
}

func intToStringArr(in []int) string {
	strs := make([]string, len(in))
	for i, v := range in {
		strs[i] = strconv.Itoa(v)
	}
	result := strings.Join(strs, "\n")
	return result + "\n"
}

func (e HostImpl) Handle(req handler.HandlerRequest) handler.HandlerResponse {
	from, _ := strconv.Atoi(req.QueryParams[0].F1)
	to, _ := strconv.Atoi(req.QueryParams[1].F1)
	fibNumbers := fibonacciRange(from, to)
	return handler.HandlerResponse{Status: 200, Body: intToStringArr(fibNumbers)}
}

//go:generate wit-bindgen tiny-go ./wit --out-dir=handler
func main() {}
