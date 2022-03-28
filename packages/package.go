package main

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	// day := d.Day()
	// fmt.Printf("day = %v\n", day)

	fmt.Println(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)

	// for _, v := range []any{"hi", 42, func() {}} {
	// 	switch v := reflect.ValueOf(v); v.Kind() {
	// 	case reflect.String:
	// 		fmt.Println(v.String())
	// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	// 		fmt.Println(v.Int())
	// 	default:
	// 		fmt.Printf("unhandled kind %s", v.Kind())
	// 	}
	// }

	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
}
