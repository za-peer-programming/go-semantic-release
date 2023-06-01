package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Message(flag.Args()...))
}

func Message(who ...string) string {
	switch len(who) {
	case 0:
		return "Hello there!"
	case 1:
		return fmt.Sprintf("Hello %s", who[0])
	default:
		return fmt.Sprintf("Hello %s", strings.Join(who, ", "))
	}
}
