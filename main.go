// Description: lolcat but using Golang
// References:
//	https://flaviocopes.com/go-tutorial-lolcat/
//	https://github.com/dmgk/faker
//		run 'go get -u syreclabs.com/go/faker' to install
//		this just generates random text for us to use
// 	https://en.wikipedia.org/wiki/ANSI_escape_code
//		to understand how we color the text

package main

import (
	"fmt"
	"math"
	"strings"

	"syreclabs.com/go/faker"
)

// this function generates the rgb rainbow
func rgb(i int) (int, int, int) {
	var f = 0.1
	ifloat := float64(i)
	return int(math.Sin(f*ifloat+0)*127 + 128),
		int(math.Sin(f*ifloat+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*ifloat+4*math.Pi/3)*127 + 128)
}

func main() {
	var phrases []string

	// this generates all of the text
	for i := 1; i < 3; i++ {
		// append is a variadic function (meaning the number of parameters can change)
		// 	its method header prob looks like (s[]T, elems ...T)
		// to directly pass a slice to a variadic function, use the ... notation
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	// join strings in slice into 1
	output := strings.Join(phrases[:], "; ")

	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)

		// we print each character in rgb rainbow
		// \033 is the ASCII escape character
		// begin ANSI Escape sequence
		// [ - control sequence introducer
		// 38 - set foreground color
		// 		is followed by 2;r;g;b
		// m - sets the appearance of the following characters
		// 0 - reset / set all attributes back to normal
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}
