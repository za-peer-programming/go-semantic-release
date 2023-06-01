package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(flag.Args()) > 4 {
		fmt.Println("You are not allowed to say hello to more than 4 people")
		os.Exit(1)
	}
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
