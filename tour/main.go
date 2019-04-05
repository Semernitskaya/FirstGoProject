package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("My favorite number is")
	fmt.Printf("%v", Pic(10, 5))
	fmt.Print(WordCount("a a bb c bb"))

	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println(NewtonGuess1(3))
	fmt.Println(NewtonGuess1(121))
	fmt.Println(NewtonGuess1(1000000))

	fmt.Println("\n", NewtonGuessHalf(3))
	fmt.Println(NewtonGuessHalf(121))
	fmt.Println(NewtonGuessHalf(1000000))

	fmt.Println("\n", NewtonGuessConvergence01(3))
	fmt.Println(NewtonGuessConvergence01(121))
	fmt.Println(NewtonGuessConvergence01(1000000))

	fmt.Println("\n", NewtonGuessConvergence001(3))
	fmt.Println(NewtonGuessConvergence001(121))
	fmt.Println(NewtonGuessConvergence001(1000000))

	fmt.Println(NewtonWithError(3))
	fmt.Println(NewtonWithError(-3))
	_, err := NewtonWithError(-3)
	if err != nil {
		fmt.Println("oh, panic!")
	}

	var b = make([]byte, 10)
	len, error := MyReader{}.Read(b)
	fmt.Println(len, error, b)

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type IPAddr [4]byte

func (p IPAddr) String() string {
	var res string
	for i, v := range p {
		res += strconv.Itoa(int(v))
		if i < cap(p)-1 {
			res += "."
		}
	}
	return res
}

func Fibonacci() func() int {
	var first, second = 0, 1
	return func() int {
		result := first
		first = second
		second += result
		return result
	}

}

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	fields := strings.Fields(s)
	for _, value := range fields {
		elem := m[value]
		m[value] = elem + 1
	}
	return m
}
