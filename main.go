package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	opts, _ := parseOpts(args)
	// res := strings.TrimLeftFunc("--c", func(r rune) bool {
	// 	fmt.Println(unicode.IsSymbol(r))
	// 	return unicode.IsMark(r)
	// })

	fmt.Println(strings.Replace("--c", "-", "", 0))
	flag.Parse()
	fmt.Println(opts)
}

// -c, --c
func parseOpts(args []string) ([]string, error) {
	opts := make([]string, 0)
	for i := 0; i < len(args); i++ {
		// --c
		arg := args[i]
		// "--c" -> true
		isOpt := arg[0] == '-'
		if isOpt {
			numMinuses := 1
			// -'-'c
			if arg[1] == '-' {
				numMinuses++
				// bad option
				if len(arg[numMinuses:]) < 2 {
					return nil, errors.New(fmt.Sprintf("invalid option at %s", arg))
				}
			}

			opts = append(opts, args[i])
		}

	}
	return opts, nil
}
