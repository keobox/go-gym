package main

/* line filters are common command that use pipes and process output
   from stdin */

/* usage

echo 'hello'   > /tmp/lines
echo 'filter' >> /tmp/lines

cat /tmp/lines | go run line-filters.go

*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// buffered scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		// buffered scanner, by default advances every new line
		// in jargon returns the current token
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
