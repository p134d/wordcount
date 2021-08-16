package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	var str string = os.Args[1]
	fmt.Println(len(strings.Fields(str)))
}
